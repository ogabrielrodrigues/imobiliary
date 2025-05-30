package router

import (
	"imobiliary/internal/api/maker"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"net/http"

	"github.com/sirupsen/logrus"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request) *httperr.HttpError

func makeHandler(handler RouteHandler, logger *logrus.Entry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			logger.Error(err.Message, logrus.WithFields(logrus.Fields{
				"method":  r.Method,
				"path":    r.URL.Path,
				"status":  err.Code,
				"message": err.Message,
			}))
			response.Json(w, err.HttpCode, err)
		}
	}
}

func setupRoutes(h *Handler, router *http.ServeMux) error {
	mh, err := maker.MakeManagerHandler(h.pool, h.config)
	if err != nil {
		return err
	}

	router.Handle("GET /manager/{manager_id}", makeHandler(mh.FindByID, h.logger))
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
