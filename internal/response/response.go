package response

import (
	"encoding/json"
	"net/http"
)

func Json(w *http.ResponseWriter, code int, data interface{}) {
	(*w).Header().Add("content-type", "application/json")

	(*w).WriteHeader(code)
	json.NewEncoder(*w).Encode(data)
}
