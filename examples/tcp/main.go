package main

import (
	"log"
	"net"

	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	jsonrpc.Register("sayHello", jsonrpc.HandlerFunc(sayHello))
	jsonrpc.Register("numPlus", jsonrpc.HandlerFunc(numPlus))

	// 设置服务器监听的地址和端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("监听失败:", err)
	}
	defer listener.Close()
	log.Println("TCP服务器正在监听 8080 端口...")

	for {
		// 接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接受连接失败:", err)
			continue
		}

		log.Println("建立连接成功, client addr: %s", conn.RemoteAddr().String())

		// 为每个连接创建一个goroutine来处理
		go jsonrpc.TcpHandler(conn)
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
