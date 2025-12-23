package router

import (
	"imobiliary/internal/api/maker"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"net/http"

	"github.com/phuslu/log"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request) *httperr.HttpError

func makeHandler(handler RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			log.Error().Msgf("%s %s [%d]. err=%v",
				r.Method,
				r.URL.Path,
				err.HttpCode,
				err,
			)

			response.Json(w, err.HttpCode, err)
		}
	}
}

func setupRoutes(h *Handler, router *http.ServeMux) {
	mh := maker.MakeManagerHandler(h.pool)

	router.Handle("GET /manager", middleware.AuthMiddleware(makeHandler(mh.FindByID)))
	router.Handle("POST /manager", makeHandler(mh.Create))
	router.Handle("POST /auth", makeHandler(mh.Authenticate))
	router.Handle("PUT /manager/reset", makeHandler(mh.ResetPassword))

	oh := maker.MakeOwnerHandler(h.pool)

	router.Handle("GET /owner/{owner_id}", middleware.AuthMiddleware(makeHandler(oh.FindByID)))
	router.Handle("POST /owner", middleware.AuthMiddleware(makeHandler(oh.Create)))
	router.Handle("GET /owner", middleware.AuthMiddleware(makeHandler(oh.FindAll)))

	ph := maker.MakePropertyHandler(h.pool)

	router.Handle("GET /property/{property_id}", middleware.AuthMiddleware(makeHandler(ph.FindByID)))
	router.Handle("POST /property", middleware.AuthMiddleware(makeHandler(ph.Create)))
	router.Handle("GET /property", middleware.AuthMiddleware(makeHandler(ph.FindAll)))

	th := maker.MakeTenantHandler(h.pool)

	router.Handle("GET /tenant/{tenant_id}", middleware.AuthMiddleware(makeHandler(th.FindByID)))
	router.Handle("POST /tenant", middleware.AuthMiddleware(makeHandler(th.Create)))
	router.Handle("GET /tenant", middleware.AuthMiddleware(makeHandler(th.FindAll)))
}
