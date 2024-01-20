package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	_ "github.com/a-h/templ"
	"log/slog"
)

type CountStore struct {
	db *sqlx.DB
}

func NewCountStore() (s *CountStore, err error) {
	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		slog.Error(err.Error())
	}

	sqlxDb := sqlx.NewDb(db, "sqlite3")

	return &CountStore{
		db: sqlxDb,
	}, nil
}

func (s CountStore) Increment(ctx context.Context) (int, error) {
	fmt.Println("INCREMENT IN DATABASE GATERWAY")

	return 0, errors.New("hello")
}
