package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	store "github.com/ogabrielrodrigues/imobiliary/internal/store/postgres"
)

func createPlans(ctx context.Context, pool *pgxpool.Pool) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		os.Exit(1)
	}

	_, err = tx.Exec(ctx, `INSERT INTO "plan" (kind, price) VALUES ('free', 0.00)`)
	if err != nil {
		tx.Rollback(ctx)
		os.Exit(1)
	}

	_, err = tx.Exec(ctx, `INSERT INTO "plan" (kind, price) VALUES ('pro', 15.99)`)
	if err != nil {
		tx.Rollback(ctx)
		os.Exit(1)
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		os.Exit(1)
	}
}

func main() {
	environment.Load()
	env := environment.Environment

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, store.PostgresConnectionString(*env))
	if err != nil {
		os.Exit(1)
	}
	defer pool.Close()

	createPlans(ctx, pool)
}
