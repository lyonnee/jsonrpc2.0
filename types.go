package jsonrpc

import (
	"encoding/json"

	"github.com/lyonnee/hvalid"
)

type Request struct {
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      int             `json:"id,omitempty"`
	JSONRPC string          `json:"jsonrpc"`
}

type Response struct {
	ID      int    `json:"id,omitempty"`
	JSONRPC string `json:"jsonrpc"`
	Result  any    `json:"result,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func GetParams[T any](req *Request, validators ...hvalid.Validator[T]) (T, error) {
	var t T
	if err := json.Unmarshal(req.Params, &t); err != nil {
		return t, err
	}

	err := hvalid.Validate(t, validators...)
	return t, err
}
