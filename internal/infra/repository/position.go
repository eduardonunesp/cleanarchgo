package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PositionRepositoryPG struct {
	conn *pgx.Conn
}

func NewPositionRepositoryPG(connStr string) *PositionRepositoryPG {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return &PositionRepositoryPG{conn}
}
