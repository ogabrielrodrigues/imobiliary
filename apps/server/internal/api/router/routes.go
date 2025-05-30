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
	mh, err := maker.MakeManagerHandler(h.pool, h.config)
	if err != nil {
		return err
	}

	//{manager_id}
	router.Handle("GET /manager", middleware.AuthMiddleware(makeHandler(mh.FindByID, h.logger), h.config.GetJwtSecret()))
	router.Handle("POST /manager", makeHandler(mh.Create, h.logger))
	router.Handle("POST /auth", makeHandler(mh.Authenticate, h.logger))

	// property_handler := factory.NewPropertyHandlerFactory(pool)

	// mux.Handle("GET /properties", middleware.AuthMiddleware(makeHandler(property_handler.FindAllByUserID)))
	// mux.Handle("GET /properties/{property_id}", middleware.AuthMiddleware(makeHandler(property_handler.FindByID)))
	// mux.Handle("POST /properties", middleware.AuthMiddleware(makeHandler(property_handler.Create)))

	// owner_handler := factory.NewOwnerHandlerFactory(pool)

	// mux.Handle("GET /owners/{owner_id}", middleware.AuthMiddleware(makeHandler(owner_handler.FindByID)))
	// mux.Handle("GET /owners", middleware.AuthMiddleware(makeHandler(owner_handler.FindAllByManagerID)))
	// mux.Handle("POST /owners", middleware.AuthMiddleware(makeHandler(owner_handler.Create)))
	// mux.Handle("PUT /owners/assign", middleware.AuthMiddleware(makeHandler(owner_handler.AssignOwnerToProperty)))
	return nil
}
