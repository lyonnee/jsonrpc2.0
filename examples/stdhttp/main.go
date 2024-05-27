package main

import (
	"log"
	"net/http"

	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	jsonrpc.Register("sayHello", jsonrpc.HandlerFunc(sayHello))
	jsonrpc.Register("numPlus", jsonrpc.HandlerFunc(numPlus))

	// 设置路由，即当访问'/'路径时调用myHandler函数
	http.HandleFunc("/rpc", jsonrpc.StdHttpHandler)

	// 启动服务器监听8080端口
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
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
