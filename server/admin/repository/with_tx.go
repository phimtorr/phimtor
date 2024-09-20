package repository

import (
	"context"
	"database/sql"

	"github.com/phimtorr/phimtor/server/repository"
)

func withTx(ctx context.Context, db *sql.DB, fn func(context.Context, *sql.Tx) error) error {
	return repository.WithTx(ctx, db, func(tx *sql.Tx) error {
		return fn(ctx, tx)
	})

}
