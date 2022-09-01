package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
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
	return &Database{conn: conn, logger: logger}, nil
}
