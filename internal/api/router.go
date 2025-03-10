package api

import (
	"net/http"
)

func Register(h *Handler, mux *http.ServeMux) {
	registerTenant(h, mux)
}

func registerTenant(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /tenant/{tenant_id}", h.GetTenant)
	mux.HandleFunc("GET /tenant", h.GetTenants)
	mux.HandleFunc("POST /tenant", h.InsertTenant)
	mux.HandleFunc("PUT /tenant", h.UpdateTenant)
	mux.HandleFunc("DELETE /tenant/{tenant_id}", h.DeleteTenant)
}
