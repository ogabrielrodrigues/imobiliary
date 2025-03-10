package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/api"
	"github.com/ogabrielrodrigues/imobiliary/internal/store/pg"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func main() {
	env := environment.LoadServerEnvironment()

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, util.ConnStr(env))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pg.New(pool))

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
