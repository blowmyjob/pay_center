package WxPayGateway

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

type AppApiService services.Service

// CloseOrder 关闭订单
//
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
func (a *AppApiService) CloseOrder(ctx context.Context, req app.CloseOrderRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in CloseOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}/close"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &app.CloseRequest{
		Mchid: req.Mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Prepay APP支付下单
//
// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
func (a *AppApiService) Prepay(ctx context.Context, req app.PrepayRequest) (resp *app.PrepayResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/app"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(app.PrepayResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
