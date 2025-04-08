package shared

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func CheckUniqueConstraint(err error) bool {
	var pg_err *pgconn.PgError
	if !errors.As(err, &pg_err) {
		return false
	}

	if pg_err.Code == "23505" {
		return true
	}

	return false
}
