package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	Pool *pgxpool.Pool
}

func New() (*PostgresStorage, error) {
	url := "postgres://illia:password@localhost:5432/db?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		return nil, fmt.Errorf("unable to connect: %w", err)
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &PostgresStorage{pool}, nil
}

func (s *PostgresStorage) Migrate() error {
	_, err := s.Pool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			body TEXT,
			status TEXT NOT NULL,
			type TEXT NOT NULL,
			created_at Date NOT NULL
		);`,
	)

	return err
}
