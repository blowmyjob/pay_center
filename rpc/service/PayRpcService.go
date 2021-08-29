package service

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/shopspring/decimal"
	"pay_center/model"
	rpc_proto "pay_center/rpc/proto"
	"pay_center/service"
)

var payService = new(service.PayService)

type PayServer struct {
	Server *xgrpc.Server
}

func (PayServer PayServer) CreatePayRecord(context context.Context, request *rpc_proto.PayRecordCreateReq) (*rpc_proto.PayRecordCreateResp, error) {
	payRecord := model.PayRecord{
		AppId:     request.AppId,
		ProductId: request.ProductId,
		UserId:    request.UserId,
		PayAmount: decimal.NewFromFloat32(request.PayAmount),
		PayMethod: request.PayMethod,
	}
	reqNo, err := payService.InsertPayOrder(payRecord)
	return &rpc_proto.PayRecordCreateResp{
		SeqNo: reqNo,
	}, err
}

func (PayServer PayServer) SelectPayRecord(context context.Context, request *rpc_proto.PayRecordSelectReq) (resp *rpc_proto.PayRecordSelectResp, err error) {
	payRecord, err := payService.SelectPayRecord(request.SeqNo)
	if err != nil {
		return resp, err
	}
	resp = &rpc_proto.PayRecordSelectResp{
		AppId:     payRecord.AppId,
		ProductId: payRecord.ProductId,
		UserId:    payRecord.UserId,
		PayAmount: payRecord.PayAmount.String(),
		PayMethod: payRecord.PayMethod,
		AppPayId:  payRecord.AppId,
	}
	return resp, err
}
