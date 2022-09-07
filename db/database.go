package db

import (
	"context"
	"database/sql"
	"fmt"

	"elasticsearch-capable-web-app/models"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Database struct {
	conn   *sql.DB
	logger zerolog.Logger
}

type ConnConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func New(cfg ConnConfig, logger zerolog.Logger) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DbName,
	)
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
