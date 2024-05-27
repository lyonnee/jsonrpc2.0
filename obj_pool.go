package jsonrpc

import "github.com/lyonnee/hpool"

var req_pool = hpool.New[Request](func() Request {
	return Request{}
})
var resp_pool = hpool.New[Response](func() Response {
	return Response{
		JSONRPC: "2.0",
	}
})

func NewRequest() Request {
	return req_pool.Get()
}

func NewResponse() Response {
	return resp_pool.Get()
}

func DropRequest(req *Request) {
	req.ID = 0
	req.Params = nil
	req.Method = ""

	req_pool.Put(*req)
}

func DropResponse(resp *Response) {
	resp.ID = 0
	resp.Error = nil
	resp.Result = nil

	resp_pool.Put(*resp)
}
