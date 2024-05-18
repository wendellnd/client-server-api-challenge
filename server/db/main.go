package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const timeout = 10 * time.Millisecond

func InsertExchange(ctx context.Context, exchange string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	db, err := sql.Open("sqlite3", "./db/db_exchange.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO tb_exchange(exchange) VALUES (?)`
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.ExecContext(ctx, exchange)
	return err
}
