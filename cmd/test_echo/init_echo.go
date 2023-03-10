package test_echo

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoWeb *echo.Echo
var visitTotal int // 统计访问量
// 自定义中间件，下面是统计访问量
func CountMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		visitTotal++
		context.Response().Header().Add("visit-total", fmt.Sprintf("%d", visitTotal))
		return next(context)
	}
}

// 初始化业务需要基础操作，下面是初始化echo，以及中间件
func init() {
	echoWeb = echo.New()
	echoWeb.Debug = true
	echoWeb.Use(middleware.Recover()) // 主要用于拦截panic错误并且在控制台打印错误日志，避免echo程序直接崩溃
	echoWeb.Use(middleware.Logger())  // Logger中间件主要用于打印http请求日志
	echoWeb.Use(CountMiddleware)
}

// 全局使用echo对象
func GetEcho() *echo.Echo {
	return echoWeb
}
