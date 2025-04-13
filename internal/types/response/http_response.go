package response

import (
	"encoding/json"
	"net/http"
)

type Response map[string]any

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

func End(w http.ResponseWriter, code int, data any) {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
