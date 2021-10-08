package common

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Build() *Response {
	return &Response{
		Code:    200,
		Message: "ok",
		Data:    nil,
	}
}

func BuildData(data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: "ok",
		Data:    data,
	}
}

func BuildResponse(code int32, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
