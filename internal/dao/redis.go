package dao

import (
	"Goose/pkg/setting"
	"github.com/garyburd/redigo/redis"
)

type Pool struct {
	Pool *redis.Pool
}

func NewPool(pool *redis.Pool) *Pool {
	return &Pool{pool}
}

//NewRedisEngine 初始化数据库驱动
func NewRedisEngine(redisPoolSetting *setting.RedisPoolSettingS) (*redis.Pool, error) {
	pool := &redis.Pool{ //实例化一个连接池
		MaxIdle:     redisPoolSetting.MaxIdle,
		MaxActive:   redisPoolSetting.MaxActive,
		IdleTimeout: redisPoolSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(redisPoolSetting.Protocol, redisPoolSetting.Host)
		},
	}

	return pool, nil
}
