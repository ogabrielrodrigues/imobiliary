package main

import (
	"errors"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> cowork/api
	"net/http"
	"os"
	"os/signal"

<<<<<<< HEAD
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/api"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/entity"
)

func main() {
	user, _ := entity.NewUser(
		uuid.New(),
		"98767-F",
		"Gabriel Rodrigues",
		"gabriel.rodrigues@crecisp.gov.br",
	)

	user.GenerateAccessCode()

	fmt.Println()
	env := environment.LoadAPIEnvironment()
=======
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	api "github.com/ogabrielrodrigues/imobiliary/internal"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
)

func main() {
	env := environment.Load()
>>>>>>> cowork/api

	// ctx := context.Background()

	// pool, err := pgxpool.New(ctx, shared.ConnStr(env))
	// if err != nil {
	// 	logger.Error(logger.ErrDatabaseConnection, "err", err)
	// 	os.Exit(1)
	// }

	// defer pool.Close()

	// if err := pool.Ping(ctx); err != nil {
	// 	logger.Error(logger.ErrDatabaseConnection, "err", err)
	// 	os.Exit(1)
	// }

	handler := api.NewHandler(nil)

	go func() {
<<<<<<< HEAD
		if err := http.ListenAndServe(env.SERVER_ADDR, handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
=======
		logger.Info("Starting server on port " + env.SERVER_ADDR)
		if err := http.ListenAndServe(env.SERVER_ADDR, middleware.CORSMiddleware(env, handler)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Error("err", err)
				os.Exit(1)
>>>>>>> cowork/api
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
