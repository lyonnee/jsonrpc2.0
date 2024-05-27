package jsonrpc

import (
	"errors"

	"github.com/lyonnee/hmap"
	"github.com/lyonnee/hvalid"
)

type Handler interface {
	Handle(req *Request, resp *Response)
}

type HandlerFunc func(req *Request, resp *Response)

func (fn HandlerFunc) Handle(req *Request, resp *Response) {
	fn(req, resp)
}

var handlers = hmap.New[string, Handler]()

func Register(name string, handler Handler) error {
	if _, ok := handlers.LoadOrStore(name, handler); ok {
		return errors.New("handler multiple registration")
	}

	return nil
}

func requestValidator() hvalid.ValidatorFunc[*Request] {
	return hvalid.ValidatorFunc[*Request](func(req *Request) error {
		if req.JSONRPC != "2.0" {
			return errors.New("jsonrpc version not match")
		}

		return nil
	})
}

func call(req *Request, resp *Response) {
	resp.ID = req.ID

	// 校验request合法
	if err := hvalid.Validate[*Request](req, requestValidator()); err != nil {
		resp.Error = InvalidRequest
		return
	}

	defer func() {
		if err := recover(); err != nil {
			resp.Error = InternalError
			recoverHandler(err)
		}
	}()

	handler, ok := handlers.Load(req.Method)
	if !ok {
		resp.Error = MethodNotFound
		return
	}

	handler.Handle(req, resp)
}
