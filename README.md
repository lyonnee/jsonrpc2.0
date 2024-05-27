<div align="center">
</br>

# jsonrpc2.0

| English | [中文](README_zh.md) |
| --- | --- |

`jsonrpc2.0` is a lightweight JSON-RPC library implemented in Go, supporting communication via HTTP and TCP. It includes both client and server implementations, supports custom exception handling, and is concurrency-safe.
</div>

## Features
- Supports JSON-RPC 2.0 protocol.
- HTTP and TCP client/server implementations.
- Custom exception handling.
- Concurrency-safe.

## Installation
Install using the `go get` command:
```bash
go get github.com/lyonnee/jsonrpc2.0
```

### Usage Examples

#### HTTP Server

```go
import (
	"log"
	"net/http"
	"github.com/lyonnee/jsonrpc2.0"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", jsonrpc.GinHandler)
	log.Fatal(r.Run(":8080"))
}
```

#### HTTP Client

```go
import (
	"fmt"
	"log"
	"github.com/lyonnee/jsonrpc2.0"
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

#### TCP Client

```go
import (
	"fmt"
	"log"
	"net"
	"github.com/lyonnee/jsonrpc2.0"
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

#### TCP Server

```go
import (
	"log"
	"net"
	"github.com/lyonnee/jsonrpc2.0"
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

#### Custom Exception Handling
You can set a custom exception handling function:
```go
jsonrpc.SetRecoverHandler(func(err any) {
    fmt.Printf("Recovered from error: %v\n", err)
})
```

## Contributing
Issues and pull requests are welcome to improve `jsonrpc2.0`.

## License
`jsonrpc2.0` is released under the MIT License. See the [LICENSE](LICENSE) file for more information.