package main

import (
	"flag"
	"fmt"
	"pay_center/config"
	"pay_center/initServer"
)

var env string

func init() {
	flag.StringVar(&env, "e", "", "config file path")
	flag.Parse()
}

// 入口函数
func main() {
	fmt.Printf("env:%v\n", env)
	initServer.ReadProperties(env)
	config.GvaDb = initServer.GormMysql()
	config.GvaRedis = initServer.InitRedisClient()
	initServer.InitHttpServer()
}
