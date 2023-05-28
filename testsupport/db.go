package testsupport

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetDb(connStr string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
