package psqlaccesscontrol

import (
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"time"
)

func scanPermission(scanner postgres.Scanner) (*entity.Permission, error) {
	var p entity.Permission
	var createdAt time.Time
	err := scanner.Scan(&p.ID, &p.Title, &createdAt)
	return &p, err
}
