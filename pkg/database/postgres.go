package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(ctx context.Context, dataSource string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dataSource)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 25
	config.MinIdleConns = 5
	config.MaxConnIdleTime = 1 * time.Minute
	config.MaxConnLifetime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	config.BeforeConnect = func(ctx context.Context, cc *pgx.ConnConfig) error {
		log.Println("New connection is being established...")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
