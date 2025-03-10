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

func (h *Handler) GetOwner(w http.ResponseWriter, r *http.Request) {
	raw_id := r.PathValue("owner_id")
	owner_id, err := uuid.Parse(raw_id)

	if err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_UUID_INVALID))
		return
	}

	ctx := context.Background()
	owner, err := h.query.GetOwner(ctx, owner_id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.Json(&w, http.StatusNotFound, rerr.ErrorResponse(rerr.ERR_NOT_FOUND))
			return
		}

		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	response.Json(&w, http.StatusOK, kind.Response{
		"owner": owner,
	})
}

func (h *Handler) GetOwners(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	owners, err := h.query.GetOwners(ctx)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	if owners == nil {
		response.Json(&w, http.StatusOK, kind.Response{
			"message": "no owners were found",
		})
		return
	}

	response.Json(&w, http.StatusOK, kind.Response{
		"owners": owners,
	})
}

func (h *Handler) InsertOwner(w http.ResponseWriter, r *http.Request) {
	var body pg.InsertOwnerParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_INVALID_BODY))
		return
	}

	ctx := context.Background()
	id, err := h.query.InsertOwner(ctx, body)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	response.Json(&w, http.StatusCreated, kind.Response{
		"owner": id,
	})
}

func (h *Handler) UpdateOwner(w http.ResponseWriter, r *http.Request) {
	var body pg.UpdateOwnerParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_INVALID_BODY))
		return
	}

	ctx := context.Background()
	err := h.query.UpdateOwner(ctx, body)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteOwner(w http.ResponseWriter, r *http.Request) {
	raw_id := r.PathValue("owner_id")
	owner_id, err := uuid.Parse(raw_id)

	if err != nil {
		response.Json(&w, http.StatusBadRequest, rerr.ErrorResponse(rerr.ERR_UUID_INVALID))
		return
	}

	ctx := context.Background()
	err = h.query.DeleteOwner(ctx, owner_id)
	if err != nil {
		response.Json(&w, http.StatusInternalServerError, rerr.ErrorResponse(rerr.ERR_INTERNAL_SERVER))
		return
	}

	w.WriteHeader(http.StatusOK)
}
