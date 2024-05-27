package jsonrpc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/lyonnee/hmap"
)

type Client interface {
	Call(req *Request) (*Response, error)
}

type HttpClient struct {
	endpoint string
	conn     *http.Client
}

func NewHttpClient(endpoint string) (*HttpClient, error) {
	return &HttpClient{
		endpoint: endpoint,
		conn:     &http.Client{},
	}, nil
}

func (cli *HttpClient) Call(req *Request) (*Response, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, _ := http.NewRequest(
		http.MethodPost,
		cli.endpoint,
		bytes.NewBuffer(r),
	)

	httpResp, err := cli.conn.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var v []byte
	if _, err := httpResp.Body.Read(v); err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var resp Response
	if err := json.Unmarshal(v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type TcpClient struct {
	conn      net.Conn
	responses *hmap.Map[int, chan *Response]
	timeout   time.Duration
}

func NewTcpClient(conn net.Conn) (*TcpClient, error) {
	cli := TcpClient{
		conn:      conn,
		responses: hmap.New[int, chan *Response](),
		timeout:   3 * time.Second,
	}

	go cli.listen()
	return &cli, nil
}

func (c *TcpClient) listen() {
	reader := bufio.NewReader(c.conn)
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

		var resp Response
		if err := json.Unmarshal(data, &resp); err != nil {
			continue
		}
		if ch, ok := c.responses.Load(resp.ID); ok {
			ch <- &resp
			c.responses.Delete(resp.ID)
		}
	}

	c.conn.Close()
}

func (cli *TcpClient) Call(req *Request) (*Response, error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *Response)
	if _, exist := cli.responses.LoadOrStore(req.ID, ch); exist {
		return nil, errors.New("duplicate request")
	}

	msg := prependLength(bs)
	if _, err := cli.conn.Write(msg); err != nil {
		cli.responses.Delete(req.ID)
		return nil, err
	}

	select {
	case resp := <-ch:
		return resp, nil
	case <-time.After(cli.timeout):
		cli.responses.Delete(req.ID)
		return nil, errors.New("request timed out")
	}
}
