package psqlaccesscontrol

import (
	"github.com/younesbeheshti/gocast_game/repository/postgres"
)

type DB struct {
	conn *postgres.PostgresDB
}

func New(conn *postgres.PostgresDB) *DB {
	return &DB{
		conn: conn,
	}
}
