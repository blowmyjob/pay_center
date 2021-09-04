package config

import (
	"github.com/douyu/jupiter/pkg/client/rocketmq"
	"github.com/douyu/jupiter/pkg/store/gorm"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

var (
	GVA_DB        *gorm.DB
	GVA_MQ        *rocketmq.Producer
	GVA_WX_CLIENT *core.Client
)
