package main

import (
	"github.com/douyu/jupiter"
	compound_registry "github.com/douyu/jupiter/pkg/registry/compound"
	"github.com/douyu/jupiter/pkg/registry/etcdv3"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	eng := NewEngine()
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
	return eng.Serve(server)
}

func (eng *Engine) serveGRPC() error {
	server := xgrpc.StdConfig("grpc").Build()
	helloworld.RegisterGreeterServer(server.Server, &Greeter{
		server: server,
	})
	return eng.Serve(server)
}

type Greeter struct {
	server *xgrpc.Server
}

func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "Hello Jupiter, I'm " + g.server.Address(),
	}, nil
}
