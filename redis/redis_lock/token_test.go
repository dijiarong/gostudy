package redis_lock

import (
	"gostudy/redis/redis_lock/utils"
	"testing"
)

func TestTokenStability(t *testing.T) {
	// 测试相同 key 生成相同 token
	token1 := utils.GenerateStableToken("test_key")
	token2 := utils.GenerateStableToken("test_key")

	if token1 != token2 {
		t.Errorf("相同 key 应该生成相同的 token, got: %s vs %s", token1, token2)
	}

	// 测试不同 key 生成不同 token
	token3 := utils.GenerateStableToken("different_key")

	if token1 == token3 {
		t.Errorf("不同 key 应该生成不同的 token, got same token: %s", token1)
	}
}

func TestRedisLockTokenConsistency(t *testing.T) {
	// 在同一个协程中创建相同 key 的多个锁实例
	lock1 := &RedisLock{
		key:   "test_key",
		token: utils.GenerateStableToken("test_key"),
	}

	lock2 := &RedisLock{
		key:   "test_key",
		token: utils.GenerateStableToken("test_key"),
	}

	lock3 := &RedisLock{
		key:   "different_key",
		token: utils.GenerateStableToken("different_key"),
	}

	// 验证相同 key 的锁有相同的 token
	if lock1.token != lock2.token {
		t.Errorf("相同 key 的锁应该有相同的 token, got: %s vs %s", lock1.token, lock2.token)
	}

	// 验证不同 key 的锁有不同的 token
	if lock1.token == lock3.token {
		t.Errorf("不同 key 的锁应该有不同的 token, got same token: %s", lock1.token)
	}

	t.Logf("Lock1 (test_key) token: %s", lock1.token)
	t.Logf("Lock2 (test_key) token: %s", lock2.token)
	t.Logf("Lock3 (different_key) token: %s", lock3.token)
}
