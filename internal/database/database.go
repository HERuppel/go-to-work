package database

import (
	"context"
	"go-to-work/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabasePool() (*pgxpool.Pool, error) {

	config, err := pgxpool.ParseConfig(config.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10
	config.ConnConfig.ConnectTimeout = 10 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
