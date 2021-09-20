package service

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"pay_center/model"
	"pay_center/service/WxPayGateway"
	"strconv"
)

var wxPayService = new(WxPayGateway.AppApiService)

type PayGateway interface {
	PayOrder(payRecord model.PayRecord, account model.PayAccount, context context.Context) (req string)
}

type WxPayment struct {
}

func (wx *WxPayment) PayOrder(record model.PayRecord, account model.PayAccount, context context.Context) (req string) {
	tradeId := strconv.FormatInt(record.Id, 10)
	url := ""
	amount := record.PayAmount.IntPart()
	request := app.PrepayRequest{
		Appid:      &record.AppId,
		Mchid:      &record.AppId,
		OutTradeNo: &tradeId,
		NotifyUrl:  &url,
		Amount: &app.Amount{
			Total: &amount,
		},
	}
	_, _, _ = wxPayService.Prepay(context, request)
	return ""
}

type AliPayment struct {
}

func (ali *AliPayment) PayOrder(record model.PayRecord, account model.PayAccount, context context.Context) (req string) {
	return ""
}

type CashContext struct {
	Strategy PayGateway
}

func NewCashContext(cashType string) CashContext {
	c := new(CashContext)
	//这里事实上是简易工厂模式的变形，用来生产策略
	switch cashType {
	case "wx":
		c.Strategy = new(WxPayment)
	case "ali":
		c.Strategy = new(AliPayment)
	}
	return *c
}
