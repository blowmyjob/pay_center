package mq

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"pay_center/config"
)

func SendMsg(msg interface{}, topic string) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, msg)
	if err != nil {

	}
	_, _ = config.GVA_MQ_PRODUCER.SendSync(context.Background(),
		primitive.NewMessage(topic, buf.Bytes()))
}
