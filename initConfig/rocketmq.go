package initConfig

import (
	"github.com/douyu/jupiter/pkg/client/rocketmq"
	"pay_center/config"
)

func initMq(name string) {
	config.GVA_MQ = rocketmq.StdProducerConfig(name).Build()
}
