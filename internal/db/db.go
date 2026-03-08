package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDB(connStr string) (*pgx.Conn, error) {

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	// connect to neon postgresql
	if err := conn.Ping(ctx); err != nil {
		conn.Close(ctx)
		return nil, err
	}

	return conn, nil
}
