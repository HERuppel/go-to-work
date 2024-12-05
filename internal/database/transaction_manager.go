package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func WithTransaction[T any](db *pgxpool.Pool, ctx context.Context, fn func(tx pgx.Tx) (T, error)) (T, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println("ERROR_STARTING_TRANSACTION:", err)
		var zero T
		return zero, err
	}

	result, err := fn(tx)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			log.Println("ROLLBACK_ERROR:", rollbackErr)
		}
		var zero T
		return zero, err
	}

	if commitErr := tx.Commit(ctx); commitErr != nil {
		log.Println("COMMIT_ERROR:", commitErr)
		var zero T
		return zero, commitErr
	}

	return result, nil
}
