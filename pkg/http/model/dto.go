package httpModel

type BaseResponse struct {
	Error   *ErrorResponse `json:"error,omitempty"`
	Success bool           `json:"success"`
	Data    any            `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
