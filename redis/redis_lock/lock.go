package redis_lock

import (
	"context"
	"errors"
	"fmt"
	"gostudy/redis/redis_lock/utils"
	"sync/atomic"
	"time"

	"github.com/gomodule/redigo/redis"
)

const RedisLockKeyPrefix = "REDIS_LOCK_PREFIX_"

var (
	ErrLockAcquiredByOthers = errors.New("lock is acquired by others")
	ErrLockNotOwnedByYou    = errors.New("can not operate lock without ownership")
	ErrNil                  = redis.ErrNil
)

func IsRetryableErr(err error) bool {
	return errors.Is(err, ErrLockAcquiredByOthers)
}

// 基于 redis 实现的分布式锁，不可重入，但保证了对称性
type RedisLock struct {
	LockOptions
	key    string
	token  string
	client LockClient

	// 看门狗运作标识
	runningDog int32
	// 停止看门狗
	stopDog context.CancelFunc
}

func NewRedisLock(key string, client LockClient, opts ...LockOption) *RedisLock {
	r := RedisLock{
		key:    key,
		token:  utils.GetProcessAndGoroutineIDStr(),
		client: client,
	}

	for _, opt := range opts {
		opt(&r.LockOptions)
	}

	repairLock(&r.LockOptions)
	return &r
}

// Lock 加锁.
func (r *RedisLock) Lock(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			return
		}
		// 加锁成功的情况下，会启动看门狗
		// 关于该锁本身是不可重入的，所以不会出现同一把锁下看门狗重复启动的情况
		r.watchDog(ctx)
	}()

	// 不管是不是阻塞模式，都要先获取一次锁
	err = r.tryLock(ctx)
	if err == nil {
		return nil
	}

	// 非阻塞模式加锁失败直接返回错误
	if !r.isBlock {
		return err
	}

	// 判断错误是否可以允许重试，不可允许的类型则直接返回错误
	if !IsRetryableErr(err) {
		return err
	}

	// 基于阻塞模式持续轮询取锁
	err = r.blockingLock(ctx)
	return
}

func (r *RedisLock) tryLock(ctx context.Context) error {
	// 首先查询锁是否属于自己
	reply, err := r.client.SetNEX(ctx, r.getLockKey(), r.token, r.expireSeconds)
	if err != nil {
		return err
	}
	if reply != 1 {
		return fmt.Errorf("reply: %d, err: %w", reply, ErrLockAcquiredByOthers)
	}

	return nil
}

// 启动看门狗
func (r *RedisLock) watchDog(ctx context.Context) {
	// 1. 非看门狗模式，不处理
	if !r.watchDogMode {
		return
	}

	// 2. 确保之前启动的看门狗已经正常回收，添加超时避免无限等待
	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	for !atomic.CompareAndSwapInt32(&r.runningDog, 0, 1) {
		select {
		case <-ctx.Done():
			return
		case <-timeout.C:
			// 超时了，可能之前的看门狗没有正常退出，强制设置状态
			atomic.StoreInt32(&r.runningDog, 1)
			break
		default:
			time.Sleep(10 * time.Millisecond) // 短暂休眠避免CPU占用过高
		}
	}

	// 3. 启动看门狗
	ctx, r.stopDog = context.WithCancel(ctx)
	go func() {
		defer func() {
			atomic.StoreInt32(&r.runningDog, 0)
		}()
		r.runWatchDog(ctx)
	}()
}

func (r *RedisLock) runWatchDog(ctx context.Context) {
	ticker := time.NewTicker(WatchDogWorkStepSeconds * time.Second)
	defer ticker.Stop()

	// 连续失败计数器
	var consecutiveFailures int
	const maxConsecutiveFailures = 3

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

		// 看门狗负责在用户未显式解锁时，持续为分布式锁进行续期
		// 通过 lua 脚本，延期之前会确保保证锁仍然属于自己
		// 为避免因为网络延迟而导致锁被提前释放的问题，watch dog 续约时需要把锁的过期时长额外增加 5 s
		err := r.DelayExpire(ctx, WatchDogWorkStepSeconds+5)
		if err != nil {
			consecutiveFailures++

			// 如果是锁不属于自己的错误，说明锁已经被释放或过期，应该停止看门狗
			if errors.Is(err, ErrLockNotOwnedByYou) {
				// 锁已经不属于自己，停止看门狗
				return
			}

			// 连续失败次数过多，停止看门狗避免无效续期
			if consecutiveFailures >= maxConsecutiveFailures {
				// 可以在这里添加日志记录
				// log.Printf("watchdog stopped due to consecutive failures: %v", err)
				return
			}

			// 网络错误等可重试的错误，继续尝试
			continue
		}

		// 续期成功，重置失败计数器
		consecutiveFailures = 0
	}
}

// 更新锁的过期时间，基于 lua 脚本实现操作原子性
func (r *RedisLock) DelayExpire(ctx context.Context, expireSeconds int64) error {
	keysAndArgs := []interface{}{r.getLockKey(), r.token, expireSeconds}
	reply, err := r.client.Eval(ctx, LuaCheckAndExpireDistributionLock, 1, keysAndArgs)
	if err != nil {
		return fmt.Errorf("failed to execute expire script: %w", err)
	}

	ret, ok := reply.(int64)
	if !ok {
		return fmt.Errorf("unexpected reply type from expire script: %T, value: %v", reply, reply)
	}

	if ret != 1 {
		return ErrLockNotOwnedByYou
	}

	return nil
}

func (r *RedisLock) blockingLock(ctx context.Context) error {
	// 阻塞模式等锁时间上限
	timeoutCh := time.After(time.Duration(r.blockWaitingSeconds) * time.Second)
	// 轮询 ticker，每隔 50 ms 尝试取锁一次
	ticker := time.NewTicker(time.Duration(50) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		select {
		// ctx 终止了
		case <-ctx.Done():
			return fmt.Errorf("lock failed, ctx timeout, err: %w", ctx.Err())
			// 阻塞等锁达到上限时间
		case <-timeoutCh:
			return fmt.Errorf("block waiting time out, err: %w", ErrLockAcquiredByOthers)
		// 放行
		default:
		}

		// 尝试取锁
		err := r.tryLock(ctx)
		if err == nil {
			// 加锁成功，返回结果
			return nil
		}

		// 不可重试类型的错误，直接返回
		if !IsRetryableErr(err) {
			return err
		}
	}

	// 不可达
	return nil
}

// Unlock 解锁. 基于 lua 脚本实现操作原子性.
func (r *RedisLock) Unlock(ctx context.Context) error {
	defer func() {
		// 停止看门狗
		if r.stopDog != nil {
			r.stopDog()
		}
	}()

	keysAndArgs := []interface{}{r.getLockKey(), r.token}
	reply, err := r.client.Eval(ctx, LuaCheckAndDeleteDistributionLock, 1, keysAndArgs)
	if err != nil {
		return fmt.Errorf("failed to execute unlock script: %w", err)
	}

	ret, ok := reply.(int64)
	if !ok {
		return fmt.Errorf("unexpected reply type from unlock script: %T, value: %v", reply, reply)
	}

	if ret != 1 {
		return ErrLockNotOwnedByYou
	}

	return nil
}

func (r *RedisLock) getLockKey() string {
	return RedisLockKeyPrefix + r.key
}
