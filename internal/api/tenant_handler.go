package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/rerr"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
	"github.com/ogabrielrodrigues/imobiliary/internal/store/pg"
)

func (h *Handler) GetTenant(w http.ResponseWriter, r *http.Request) {
	raw_id := r.PathValue("tenant_id")
	tenant_id, err := uuid.Parse(raw_id)

	if err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_UUID_INVALID))
		return
	}

	ctx := context.Background()
	tenant, err := h.query.GetTenant(ctx, tenant_id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.Json(&w, http.StatusNotFound, rerr.ErrorResponse(rerr.ERR_TENANT_NOT_FOUND))
			return
		}

		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	response.Json(&w, http.StatusOK, kind.Response{
		"tenant": tenant,
	})
}

func (h *Handler) GetTenants(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tenants, err := h.query.GetTenants(ctx)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	if tenants == nil {
		response.Json(&w, http.StatusOK, kind.Response{
			"message": "no tenants were found",
		})
		return
	}

	response.Json(&w, http.StatusOK, kind.Response{
		"tenants": tenants,
	})
}

func (h *Handler) InsertTenant(w http.ResponseWriter, r *http.Request) {
	var body pg.InsertTenantParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_INVALID_TENANT_BODY))
		return
	}

	ctx := context.Background()
	id, err := h.query.InsertTenant(ctx, body)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	response.Json(&w, http.StatusCreated, kind.Response{
		"tenant": id,
	})
}

func (h *Handler) UpdateTenant(w http.ResponseWriter, r *http.Request) {
	var body pg.UpdateTenantParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_INVALID_TENANT_BODY))
		return
	}

	ctx := context.Background()
	err := h.query.UpdateTenant(ctx, body)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteTenant(w http.ResponseWriter, r *http.Request) {
	raw_id := r.PathValue("tenant_id")
	tenant_id, err := uuid.Parse(raw_id)

	if err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_UUID_INVALID))
		return
	}

	ctx := context.Background()
	err = h.query.DeleteTenant(ctx, tenant_id)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	w.WriteHeader(http.StatusOK)
}
