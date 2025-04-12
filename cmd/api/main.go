package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

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

	// ctx := context.Background()

	// pool, err := pgxpool.New(ctx, shared.ConnStr(env))
	// if err != nil {
	// 	panic(err)
	// }

	// defer pool.Close()

	// if err := pool.Ping(ctx); err != nil {
	// 	panic(err)
	// }

	handler := api.NewHandler(nil)

	go func() {
		if err := http.ListenAndServe(env.SERVER_ADDR, handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
