package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
)

func main() {
	eng := NewEngine()
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
