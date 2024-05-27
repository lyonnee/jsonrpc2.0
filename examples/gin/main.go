package main

import (
	"github.com/gin-gonic/gin"
	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	jsonrpc.Register("sayHello", jsonrpc.HandlerFunc(sayHello))
	jsonrpc.Register("numPlus", jsonrpc.HandlerFunc(numPlus))

	// 创建一个默认的Gin路由器
	r := gin.Default()

	// 定义一个处理函数，响应GET请求
	r.POST("/rpc", jsonrpc.GinHandler)

	// 在8080端口启动服务
	r.Run(":8080")
}

func sayHello(req *jsonrpc.Request, resp *jsonrpc.Response) {
	name, err := jsonrpc.GetParams[string](req)
	if err != nil {
		resp.Error = jsonrpc.ParseError
	}

	resp.Result = ("hello " + name)
}

type Nums struct {
	A int `json:"a"`
	B int `json:"b"`
}

func numPlus(req *jsonrpc.Request, resp *jsonrpc.Response) {
	nums, err := jsonrpc.GetParams[Nums](req)
	if err != nil {
		resp.Error = jsonrpc.InvalidParams
		return
	}

	resp.Result = (nums.A + nums.B)
}
