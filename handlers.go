package jsonrpc

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"io"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinHandler(c *gin.Context) {
	var req, resp = NewRequest(), NewResponse()
	defer func() {
		// 回收req,resp对象
		DropRequest(&req)
		DropResponse(&resp)
	}()

	if err := c.ShouldBind(&req); err != nil {
		resp.Error = ParseError

		c.JSON(
			http.StatusOK,
			resp,
		)

		return
	}

	call(&req, &resp)

	if req.ID == 0 {
		return
	}

	c.JSON(
		http.StatusOK,
		resp,
	)
}

func StdHttpHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	var req, resp = NewRequest(), NewResponse()
	defer func() {
		DropRequest(&req)
		DropResponse(&resp)
	}()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Error = ParseError

		bs, _ := json.Marshal(resp)
		w.Write(bs)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		resp.Error = ParseError

		bs, _ := json.Marshal(resp)
		w.Write(bs)
		return
	}

	call(&req, &resp)
	if req.ID == 0 {
		return
	}

	res, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func TcpHandler(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		lengthBytes := make([]byte, 4)
		_, err := reader.Read(lengthBytes)
		if err != nil {
			break
		}

		length := binary.BigEndian.Uint32(lengthBytes)
		data := make([]byte, length)
		_, err = reader.Read(data)
		if err != nil {
			break
		}

		var req, resp = NewRequest(), NewResponse()
		defer func() {
			DropRequest(&req)
			DropResponse(&resp)
		}()

		if err := json.Unmarshal(data, &req); err != nil {
			resp.Error = ParseError

			bs, _ := json.Marshal(resp)
			msg := prependLength(bs)
			conn.Write(msg)

			continue
		}

		call(&req, &resp)

		if req.ID == 0 {
			continue
		}

		res, _ := json.Marshal(resp)
		msg := prependLength(res)
		if _, err := conn.Write(msg); err != nil {
			return
		}
	}

	// log.Println("连接断开")
}

// prependLength 添加消息长度前缀
func prependLength(data []byte) []byte {
	length := uint32(len(data))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)
	return append(lengthBytes, data...)
}
