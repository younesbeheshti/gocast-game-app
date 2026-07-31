package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Sslmode  string
}

type PostgresDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) PostgresDB {
	//dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.Sslmode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return PostgresDB{db: db, config: config}
}
