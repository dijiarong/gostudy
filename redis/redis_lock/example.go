package redis_lock

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ExampleUsage 展示分布式锁的使用示例
func ExampleUsage() {
	// 创建 Redis 客户端
	client := NewClient("tcp", "localhost:6379", "", WithMaxIdle(10))

	// 创建分布式锁，启用看门狗模式
	lock := NewRedisLock("my_business_key", client)

	ctx := context.Background()

	// 加锁
	err := lock.Lock(ctx)
	if err != nil {
		log.Printf("Failed to acquire lock: %v", err)
		return
	}

	fmt.Println("Lock acquired successfully")

	// 模拟业务处理
	time.Sleep(35 * time.Second) // 超过默认锁过期时间，看门狗会自动续期

	// 解锁
	err = lock.Unlock(ctx)
	if err != nil {
		log.Printf("Failed to unlock: %v", err)
		return
	}

	fmt.Println("Lock released successfully")
}

// ExampleTokenStability 展示 token 稳定性的示例
func ExampleTokenStability() {
	client := NewClient("tcp", "localhost:6379", "", WithMaxIdle(10))

	// 在同一个协程中，对相同的 key 创建多个锁实例
	lock1 := NewRedisLock("same_key", client)
	lock2 := NewRedisLock("same_key", client)
	lock3 := NewRedisLock("different_key", client)

	fmt.Printf("Lock1 token: %s\n", lock1.token)
	fmt.Printf("Lock2 token: %s\n", lock2.token)
	fmt.Printf("Lock3 token: %s\n", lock3.token)

	// lock1 和 lock2 应该有相同的 token（相同的 key）
	// lock3 应该有不同的 token（不同的 key）
	if lock1.token == lock2.token {
		fmt.Println("✅ 相同 key 的锁实例有相同的 token")
	} else {
		fmt.Println("❌ 相同 key 的锁实例应该有相同的 token")
	}

	if lock1.token != lock3.token {
		fmt.Println("✅ 不同 key 的锁实例有不同的 token")
	} else {
		fmt.Println("❌ 不同 key 的锁实例应该有不同的 token")
	}
}

// ExampleRedLockUsage 展示红锁的使用示例
func ExampleRedLockUsage() {
	// 配置多个 Redis 节点
	confs := []*SingleNodeConf{
		{Network: "tcp", Address: "localhost:6379", Password: ""},
		{Network: "tcp", Address: "localhost:6380", Password: ""},
		{Network: "tcp", Address: "localhost:6381", Password: ""},
	}

	// 创建红锁
	redLock, err := NewRedLock("my_business_key", confs,
		WithSingleNodesTimeout(100*time.Millisecond),
		WithRedLockExpireDuration(30*time.Second))
	if err != nil {
		log.Printf("Failed to create RedLock: %v", err)
		return
	}

	ctx := context.Background()

	// 加锁
	err = redLock.Lock(ctx)
	if err != nil {
		log.Printf("Failed to acquire RedLock: %v", err)
		return
	}

	fmt.Println("RedLock acquired successfully")

	// 模拟业务处理
	time.Sleep(5 * time.Second)

	// 解锁
	err = redLock.Unlock(ctx)
	if err != nil {
		log.Printf("Failed to unlock RedLock: %v", err)
		return
	}

	fmt.Println("RedLock released successfully")
}
