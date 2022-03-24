package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	DBEngine  *gorm.DB
	RedisPool *redis.Pool
)
