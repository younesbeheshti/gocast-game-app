package postgres

import (
	"database/sql"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"time"
)

func (d *PostgresDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time
	err := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("cant scan query result: %w", err)
	}
	return false, nil
}
func (d *PostgresDB) Register(u entity.User) (*entity.User, error) {
	err := d.db.QueryRow(
		`INSERT INTO users(name, phone_number)
		 VALUES($1, $2)
		 RETURNING id`,
		u.Name,
		u.PhoneNumber,
	).Scan(&u.ID)

	if err != nil {
		return nil, fmt.Errorf("can't execute command: %w", err)
	}

	return &u, nil
}
