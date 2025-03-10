package rerr

import "net/http"

type Code int

const (
	ERR_UUID_INVALID        Code = 400
	ERR_INVALID_TENANT_BODY Code = 400
	ERR_TENANT_NOT_FOUND    Code = 404
	ERR_INTERNAL_SERVER     Code = 500
)

type Err struct {
	Message string `json:"message"`
	Code    Code   `json:"code"`
}

func (e *Err) Error() string {
	return e.Message
}

func ErrorResponse(code Code) Err {
	return Err{
		Message: http.StatusText(int(code)),
		Code:    code,
	}
}
