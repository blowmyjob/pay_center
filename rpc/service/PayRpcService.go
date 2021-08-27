package service

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	rpc_proto "pay_center/rpc/proto"
)

type PayServer struct {
	Server *xgrpc.Server
}

func (PayServer PayServer) CreatePayRecord(context context.Context, request *rpc_proto.PayRecordCreateReq) (*rpc_proto.PayRecordCreateResp, error) {
	return &rpc_proto.PayRecordCreateResp{
		SeqNo: int64(1),
	}, nil
}

func (PayServer PayServer) SelectPayRecord(context context.Context, request *rpc_proto.PayRecordSelectReq) (*rpc_proto.PayRecordSelectResp, error) {
	return nil, nil
}
