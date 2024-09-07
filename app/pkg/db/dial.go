package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func Dial(ctx context.Context, url string) (*pgxpool.Pool, error) {
	fmt.Println(url)
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse DATABASE_URL: %v\n", err)
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %v\n", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to the database")
	return dbpool, nil
}
