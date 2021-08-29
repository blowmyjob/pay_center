package config

import (
	"github.com/douyu/jupiter/pkg/client/rocketmq"
	"github.com/douyu/jupiter/pkg/store/gorm"
)

var (
	GVA_DB *gorm.DB
	GVA_MQ *rocketmq.Producer
)
