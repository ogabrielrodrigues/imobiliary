package response

import (
	"encoding/json"
	"net/http"
)

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

func (e *Err) Error() string {
	return e.Message
}

func NewErr(code int, message string) *Err {
	return &Err{
		Code:    code,
		Err:     http.StatusText(code),
		Message: message,
	}
}

type response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Result any    `json:"result"`
}

func End(w http.ResponseWriter, code int, data any) *Err {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(response{
		Status: http.StatusText(code),
		Code:   code,
		Result: data,
	}); err != nil {
		return NewErr(http.StatusBadRequest, ERR_SERIALIZING_JSON)
	}

	return nil
}
