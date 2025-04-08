package response

import (
	"encoding/json"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func Json(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, code int, err error) {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(kind.Response{
		"code":    code,
		"error":   http.StatusText(code),
		"message": err.Error(),
	})
}
