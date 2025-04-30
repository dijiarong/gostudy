package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("first_mcp", "1.0.0")
	// 添加工具

	s.AddTool(mcp.NewTool("calculate",
		mcp.WithDescription("执行基本的算数运算"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("要执行的算数运算类"),
			mcp.Enum("multiply", "divide"),
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("第一个数字"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("第二个数字"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		op := request.Params.Arguments["operation"].(string)
		x := request.Params.Arguments["x"].(float64)
		y := request.Params.Arguments["y"].(float64)
		var res float64
		switch op {
		case "multiply":
			res = x * y
		case "divide":
			res = x / y
		}
		return mcp.FormatNumberResult(res), nil
	})
	// 启动服务
	if err := server.ServeStdio(s); err != nil {
		panic(err)
	}
}
