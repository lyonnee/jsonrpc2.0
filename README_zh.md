
<div align="center">
</br>

# jsonrpc2.0

| [English](README.md) | 中文 |
| --- | --- |

`jsonrpc2.0` 是一个用 Go 实现的轻量级 JSON-RPC 库，支持通过 HTTP 和 TCP 进行通信，包含客户端和服务端的实现，支持自定义异常处理和并发安全。
</div>

## 特性
- 支持 JSON-RPC 2.0 协议
- HTTP 和 TCP 客户端/服务端实现
- 自定义异常处理
- 并发安全

## 安装
使用`go get`命令安装:
```bash
go get github.com/lyonnee/jsonrpc2.0
```


### 使用示例

#### HTTP 服务端

```go
import (
	"log"
	jsonrpc "github.com/lyonnee/jsonrpc2.0"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", jsonrpc.GinHandler)
	log.Fatal(r.Run(":8080"))
}
```

#### HTTP 客户端

```go
import (
	"fmt"
	"log"
	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	client, err := jsonrpc.NewHttpClient("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error creating HTTP client: %v", err)
	}

	req := &jsonrpc.Request{
		ID:      1,
		Method:  "exampleMethod",
		JSONRPC: "2.0",
	}

	resp, err := client.Call(req)
	if err != nil {
		log.Fatalf("Error calling method: %v", err)
	}

	fmt.Printf("Response: %+v\n", resp)
}

```

#### TCP 客户端

```go
import (
	"fmt"
	"log"
	"net"
	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}

	client, err := jsonrpc.NewTcpClient(conn)
	if err != nil {
		log.Fatalf("Error creating TCP client: %v", err)
	}

	req := &jsonrpc.Request{
		ID:      1,
		Method:  "exampleMethod",
		JSONRPC: "2.0",
	}

	resp, err := client.Call(req)
	if err != nil {
		log.Fatalf("Error calling method: %v", err)
	}

	fmt.Printf("Response: %+v\n", resp)
}
```

#### TCP 服务端

```go
import (
	"log"
	"net"
	jsonrpc "github.com/lyonnee/jsonrpc2.0"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	log.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go jsonrpc.TcpHandler(conn)
	}
}
```

#### 自定义异常处理
可以设置自定义的异常处理函数：
```go
jsonrpc.SetRecoverHandler(func(err any) {
    fmt.Printf("Recovered from error: %v\n", err)
})
```

## 贡献
欢迎提交问题和拉取请求来改进`jsonrpc2.0`。

## 许可证
`jsonrpc2.0`遵循MIT许可证。查看[LICENSE](LICENSE)文件以获取更多信息。