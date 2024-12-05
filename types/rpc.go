package types

var noParams = make([]any, 0)

type Request struct {
	Id     uint32 `json:"id"`
	Method string `json:"method"`
	Params []any  `json:"params"`
}

func NewRequest(id uint32, method string, params ...any) *Request {
	if params == nil {
		params = noParams
	}

	return &Request{
		Id:     id,
		Method: method,
		Params: params,
	}
}

type Response struct {
	Id     uint32 `json:"id"`
	Result any    `json:"result"`
	Error  *Error `json:"error"`
}
