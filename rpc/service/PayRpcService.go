package service

import (
	"context"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/shopspring/decimal"
	"pay_center/model"
	"pay_center/mq"
	rpc_proto "pay_center/rpc/proto"
	"pay_center/service"
)

var payService = new(service.PayService)
var payAccountService = new(service.PayAccountService)

type PayServer struct {
	Server *xgrpc.Server
}

func (PayServer PayServer) RefundPay(ctx context.Context, req *rpc_proto.RefundReq) (*rpc_proto.RefundResp, error) {
	panic("implement me")
}

func (PayServer PayServer) ReversePay(ctx context.Context, req *rpc_proto.ReverseReq) (*rpc_proto.RefundResp, error) {
	panic("implement me")
}

func (PayServer PayServer) CreatePayRecord(context context.Context, request *rpc_proto.PayRecordCreateReq) (*rpc_proto.PayRecordCreateResp, error) {
	payAccount, err := payAccountService.SelectAccount(request.AppId)
	if err != nil {
		return nil, err
	}
	if payAccount.Status == "0" {
		return nil, err
	}
	payRecord := model.PayRecord{
		AppId:     request.AppId,
		ProductId: request.ProductId,
		UserId:    request.UserId,
		PayAmount: decimal.NewFromFloat32(request.PayAmount),
		PayMethod: request.PayMethod,
	}
	reqNo, err := payService.InsertPayOrder(payRecord)
	payment := service.NewCashContext(payRecord.PayMethod)
	mq.SendMsg(payRecord, "")
	payment.Strategy.PayOrder(payRecord, payAccount, context)
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

func (PayServer PayServer) refund(context context.Context, request *rpc_proto.RefundReq) (resp *rpc_proto.RefundResp, err error) {
	payRecord, err := payService.SelectPayRecord(request.OrderId)
	if err != nil {
		return nil, err
	}
	if &payRecord == nil {
		return nil, err
	}
	//?????????????????????????????????
	payService.UpdateOrderStatus(request.OrderId, "")
	return resp, nil
}
