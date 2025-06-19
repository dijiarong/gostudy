package utils

import (
	"crypto/md5"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetCurrentProcessID() string {
	return strconv.Itoa(os.Getpid())
}

// GetCurrentGoroutineID 获取当前的协程ID
func GetCurrentGoroutineID() string {
	buf := make([]byte, 128)
	buf = buf[:runtime.Stack(buf, false)]
	stackInfo := string(buf)
	return strings.TrimSpace(strings.Split(strings.Split(stackInfo, "[running]")[0], "goroutine")[1])
}

// GetProcessAndGoroutineIDStr 获取进程ID和协程ID的组合
// 这个标识在同一个协程中是稳定的，适合作为分布式锁的 token
func GetProcessAndGoroutineIDStr() string {
	return fmt.Sprintf("%s_%s", GetCurrentProcessID(), GetCurrentGoroutineID())
}

// GenerateStableToken 基于 key 和当前执行上下文生成稳定的 token
// 同一个协程对相同的 key 会生成相同的 token，但不同协程或不同进程会生成不同的 token
func GenerateStableToken(key string) string {
	// 使用进程ID + 协程ID + key 的组合来生成稳定的 token
	// 这样确保：
	// 1. 同一协程对相同 key 生成相同 token
	// 2. 不同协程或不同进程生成不同 token
	// 3. 相同协程对不同 key 生成不同 token
	baseInfo := fmt.Sprintf("%s_%s_%s", GetCurrentProcessID(), GetCurrentGoroutineID(), key)

	// 使用 MD5 哈希来生成固定长度的 token，避免 token 过长
	hash := md5.Sum([]byte(baseInfo))
	return fmt.Sprintf("%x", hash)
}
