package api

import (
	"net/http"
)

func Register(h *Handler, mux *http.ServeMux) {
	registerTenant(h, mux)
	registerOwner(h, mux)
	// registerAddress(h, mux)
}

func registerTenant(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /tenant/{tenant_id}", h.GetTenant)
	mux.HandleFunc("GET /tenant", h.GetTenants)
	mux.HandleFunc("POST /tenant", h.InsertTenant)
	mux.HandleFunc("PUT /tenant", h.UpdateTenant)
	mux.HandleFunc("DELETE /tenant/{tenant_id}", h.DeleteTenant)
}

func registerOwner(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /owner/{owner_id}", h.GetOwner)
	mux.HandleFunc("GET /owner", h.GetOwners)
	mux.HandleFunc("POST /owner", h.InsertOwner)
	mux.HandleFunc("PUT /owner", h.UpdateOwner)
	mux.HandleFunc("DELETE /owner/{owner_id}", h.DeleteOwner)
}

// func registerAddress(h *Handler, mux *http.ServeMux) {
// 	mux.HandleFunc("GET /address/{address_id}", h.GetAddress)
// 	mux.HandleFunc("GET /address", h.GetAddresses)
// 	mux.HandleFunc("POST /address", h.InsertAddress)
// 	mux.HandleFunc("PUT /address", h.UpdateAddress)
// 	mux.HandleFunc("DELETE /address/{address_id}", h.DeleteAddress)
// }
