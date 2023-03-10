package test_echo

import (
	echoSwagger "github.com/swaggo/echo-swagger"
)

// 设置router
func SetRouter() {
	HelloWorldHandler{}.router()
	PrintName{}.router()
	GetEcho().GET("/docs/*", echoSwagger.WrapHandler) // 引入swagger文档接口
}

// 启动router
func StartRouter() {
	SetRouter()
	GetEcho().Start(":8080")
}
