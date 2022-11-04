package db

import (
	"context"
	"database/sql"

	"elasticsearch-capable-web-app/models"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Database struct {
	conn   *sql.DB
	logger zerolog.Logger
}

func New(dsn string, logger zerolog.Logger) (*Database, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	upsertDefaultUser(conn)
	return &Database{conn: conn, logger: logger}, nil
}

var defaultUser = &models.User{ID: 1, Name: "aria"}

func upsertDefaultUser(db *sql.DB) {
	err := defaultUser.Upsert(
		context.Background(),
		db,
		false, // do nothing on conflict users.id
		[]string{models.UserColumns.ID},
		boil.None(),
		boil.Infer(),
	)
	if err != nil {
		panic(err)
	}
}
