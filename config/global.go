package config

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB_USER_NAME  string
	DB_USER_PASWD string
	DB_USER_HOST  string
	DB_USER_PORT  int64
	DB_TIMEOUT    string
	DB_NAME       string
	GvaDb         *gorm.DB

	REDIS_ARR  string
	REDIS_PASS string
	GvaRedis   *redis.Client

	MONGO_ARR  string
	MONGO_POOL int64
)
