package migrator

import (
	"database/sql"
	"fmt"
	"github.com/rubenv/sql-migrate"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
)

type Migrator struct {
	dialect    string
	dbConfig   postgres.Config
	migrations *migrate.FileMigrationSource
}

// TODO -set migration table name
// TODO -add limit to Up and Down method

func New(dbCfg postgres.Config) Migrator {

	migrations := &migrate.FileMigrationSource{
		Dir: "repository/postgres/migrations",
	}

	return Migrator{migrations: migrations, dbConfig: dbCfg, dialect: "postgres"}
}

func (m Migrator) Up() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.Database, m.dbConfig.Sslmode)

	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("failed to execute migration: %w", err))
	}

	fmt.Printf("migrated %d migrations\n", n)
}
func (m Migrator) Down() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.Database, m.dbConfig.Sslmode)

	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("failed rollback migrations: %w", err))
	}

	fmt.Printf("rollback %d migrations\n", n)
}
func (m *Migrator) Status() {
	// TODO - to add status
}
