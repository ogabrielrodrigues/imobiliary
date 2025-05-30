package response

import (
	"context"
	"encoding/json"
	"imobiliary/internal/application/httperr"
	"net/http"
)

type response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Result any    `json:"result"`
}

func End(w http.ResponseWriter, code int, data any) *httperr.HttpError {
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(response{
		Status: http.StatusText(code),
		Code:   code,
		Result: data,
	}); err != nil {
		return httperr.NewInternalServerError(context.Background(), "error encoding json response")
	}

	return nil
}
