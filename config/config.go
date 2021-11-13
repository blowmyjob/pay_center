package config

import (
	"github.com/douyu/jupiter/pkg/client/redis"
	"github.com/douyu/jupiter/pkg/store/gorm"
)

var (
	GVA_DB       *gorm.DB
	REDIS_CLIENT *redis.Redis
)
