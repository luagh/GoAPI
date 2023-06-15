package cache

import (
	"GOHUB/pkg/config"
	"GOHUB/pkg/redis"
	"time"
)

// RedisStore 实现 cache.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// /创建一个新的 Redis 数据存储对象，并设置连接参数和缓存键前缀
func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(address, username, password, db)
	rs.KeyPrefix = config.GetString("app.name") + ":cache:"
	return rs
}

// 、、指定键值对存储到 Redis 中，并设置过期时间
func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

// 从 Redis 中获取指定键的值
func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

// 检查指定键是否存在于 Redis 中
func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

// 从 Redis 中删除指定键及其对应的值
func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

// 将指定键值对永久存储在 Redis 中
func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

// 清空当前 Redis 数据库中的所有键值对
func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

// 指定键所对应的数值进行加一操作
func (s *RedisStore) Increment(parameters ...interface{}) {
	s.RedisClient.Increment(parameters...)
}

// 指定键所对应的数值进行减一操作
func (s *RedisStore) Decrement(parameters ...interface{}) {
	s.RedisClient.Decrement(parameters...)
}

// 检测 Redis 数据库是否正常连接，返回错误信息
func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}
