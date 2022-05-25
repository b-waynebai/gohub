package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPreFix   string
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) bool {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 	本地环境调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPreFix+key, value, ExpireTime)
}

// Get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clean bool) (value string) {
	key = s.KeyPreFix + key
	val := s.RedisClient.Get(key)
	if clean {
		s.RedisClient.Del(key)
	}

	return val
}

// Verify 实现 verifycode.Store interface 的 Verify 方法
func (s *RedisStore) Verify(id, answer string, clean bool) bool {
	v := s.Get(id, clean)

	return v == answer
}
