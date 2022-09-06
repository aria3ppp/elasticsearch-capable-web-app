package main

import (
	"database/sql"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

//go:generate sqlboiler --wipe psql

type Config struct {
	Psql struct {
		DBName   string `toml:"dbname" env-required:"true"`
		Host     string `toml:"host" env-required:"true"`
		Port     uint16 `toml:"port" env-default:"5432"`
		User     string `toml:"user" env-required:"true"`
		Password string `toml:"pass" env-required:"true"`
		SSLMode  string `toml:"sslmode" env-default:"disable"`
	} `yaml:"psql" env-required:"true"`
}

func sqlBoilerDemo() {
	// db := initDB()

	// post := &models.PostsStore{Title: "LIFE", Body: "This is about life..."}
}

func initDB() *sql.DB {
	var cfg Config
	err := cleanenv.ReadConfig("sqlboiler.toml", &cfg)
	panicIF(err)

	connString := fmt.Sprintf(
		"dbname=%s host=%s port=%d user=%s password=%s sslmode=%s",
		cfg.Psql.DBName,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.SSLMode,
	)
	db, err := sql.Open("postgres", connString)
	panicIF(err)

	err = db.Ping()
	panicIF(err)

	fmt.Println("connected...")

	return db
}

func panicIF(err error) {
	if err != nil {
		panic(err)
	}
}
