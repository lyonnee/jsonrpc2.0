package jsonrpc

const (
	// Invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	PARSE_ERR_CODE int    = -32700
	PARSE_ERR_MSG  string = "Parse error"

	// The JSON sent is not a valid Request object.
	INVALID_REQUEST_CODE int    = -32600
	INVALID_REQUEST_MSG  string = "Invalid request"

	// The method does not exist / is not available.
	METHOD_NOT_FOUND_CODE int    = -32601
	METHOD_NOT_FOUND_MSG  string = "Method not found"

	// Invalid method parameter(s).
	INVALID_PARAMS_CODE int    = -32602
	INVALID_PARAMS_MSG  string = "Invalid params"

	// Internal JSON-RPC error.
	INTERNAL_ERROR_CODE int    = -32603
	INTERNAL_ERROR_MSG  string = "Internal error"

	// Reserved for implementation-defined server-errors.
	// Code in -32000 to -32099
	SERVER_ERROR_MSG string = "Server error"
)

var (
	ParseError = Error{
		Code: PARSE_ERR_CODE,
		Msg:  PARSE_ERR_MSG,
	}

	InvalidRequest = Error{
		Code: INVALID_REQUEST_CODE,
		Msg:  INVALID_REQUEST_MSG,
	}

	MethodNotFound = Error{
		Code: METHOD_NOT_FOUND_CODE,
		Msg:  METHOD_NOT_FOUND_MSG,
	}

	InvalidParams = Error{
		Code: INVALID_PARAMS_CODE,
		Msg:  INVALID_PARAMS_MSG,
	}

	InternalError = Error{
		Code: INTERNAL_ERROR_CODE,
		Msg:  INTERNAL_ERROR_MSG,
	}
)
