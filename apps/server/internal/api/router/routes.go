package router

import (
	"imobiliary/internal/api/maker"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"net/http"

	"go.uber.org/zap"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request) *httperr.HttpError

func makeHandler(handler RouteHandler, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			logger.Error(err.Message,
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", err.HttpCode),
				zap.String("message", err.Message),
			)
			response.Json(w, err.HttpCode, err)
		}
	}
}

func setupRoutes(h *Handler, router *http.ServeMux) error {
	mh := maker.MakeManagerHandler(h.pool, h.config)

	router.Handle("GET /manager", middleware.AuthMiddleware(makeHandler(mh.FindByID, h.logger), h.config.GetJwtSecret()))
	router.Handle("POST /manager", makeHandler(mh.Create, h.logger))
	router.Handle("POST /auth", makeHandler(mh.Authenticate, h.logger))

	oh := maker.MakeOwnerHandler(h.pool, h.config)

	router.Handle("GET /owner/{owner_id}", middleware.AuthMiddleware(makeHandler(oh.FindByID, h.logger), h.config.GetJwtSecret()))
	router.Handle("POST /owner", middleware.AuthMiddleware(makeHandler(oh.Create, h.logger), h.config.GetJwtSecret()))
	router.Handle("GET /owner", middleware.AuthMiddleware(makeHandler(oh.FindAll, h.logger), h.config.GetJwtSecret()))

	ph := maker.MakePropertyHandler(h.pool, h.config)

	router.Handle("GET /property/{property_id}", middleware.AuthMiddleware(makeHandler(ph.FindByID, h.logger), h.config.GetJwtSecret()))
	router.Handle("POST /property", middleware.AuthMiddleware(makeHandler(ph.Create, h.logger), h.config.GetJwtSecret()))
	router.Handle("GET /property", middleware.AuthMiddleware(makeHandler(ph.FindAll, h.logger), h.config.GetJwtSecret()))

	th := maker.MakeTenantHandler(h.pool, h.config)

	router.Handle("GET /tenant/{tenant_id}", middleware.AuthMiddleware(makeHandler(th.FindByID, h.logger), h.config.GetJwtSecret()))
	router.Handle("POST /tenant", middleware.AuthMiddleware(makeHandler(th.Create, h.logger), h.config.GetJwtSecret()))
	router.Handle("GET /tenant", middleware.AuthMiddleware(makeHandler(th.FindAll, h.logger), h.config.GetJwtSecret()))

	return nil
}
