package mq

import (
	"bytes"
	"encoding/binary"
)

func SendMsg(msg interface{}, topic string) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, msg)
	if err != nil {

	}
}
