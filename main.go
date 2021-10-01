package main

import (
	"context"
	"fmt"
	"github.com/douyu/jupiter"
	compound_registry "github.com/douyu/jupiter/pkg/registry/compound"
	"github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"

	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/reflection"
	"pay_center/config"
	"pay_center/initConfig"
	rpc_proto "pay_center/rpc/proto"
	rpc_service "pay_center/rpc/service"
)

func main() {
	eng := NewEngine()
	config.GVA_DB = initConfig.GormMysql("test")
	eng.SetRegistry(compound_registry.New(
		etcdv3.StdConfig("wh").Build(),
	))
	if err := eng.Run(); err != nil {
		xlog.Panic(err.Error())
	}
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.serveHTTP,
		eng.serveGRPC,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

// HTTP地址
func (eng *Engine) serveHTTP() error {
	server := xgin.StdConfig("http").Build()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello Gin")
	})
	//微信支付回调
	server.GET("/wxPay/call", func(ctx *gin.Context) {
		context := context.Background()
		handler := new(notify.Handler)
		transaction := new(payments.Transaction)
		notifyReq, err := handler.ParseNotifyRequest(context, ctx.Request, transaction)
		if err != nil || notifyReq == nil {

		}
		fmt.Println(notifyReq.Summary)
		fmt.Println(transaction.OutTradeNo)
	})

	return eng.Serve(server)
}

func (eng *Engine) serveGRPC() error {
	server := xgrpc.StdConfig("grpc").Build()
	reflection.Register(server.Server)
	rpc_proto.RegisterPayServiceServer(server.Server, &rpc_service.PayServer{
		Server: server,
	})
	return eng.Serve(server)
}
