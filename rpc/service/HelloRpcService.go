package service

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	rpc_proto "pay_center/rpc/proto"
)

type Greeter struct {
	Server *xgrpc.Server
}

func (g Greeter) SayHello(context context.Context, request *rpc_proto.SayHelloReq) (*rpc_proto.SayHelloRes, error) {
	return &rpc_proto.SayHelloRes{
		Resp: "Hello Jupiter, I'm " + request.Name,
	}, nil
}
