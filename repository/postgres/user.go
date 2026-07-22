package postgres

import (
	"database/sql"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"time"
)

func (d *PostgresDB) GetUserByPhoneNumber(phoneNumber string) (*entity.User, bool, error) {
	user := &entity.User{}
	var createdAt time.Time
	err := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("cant scan query result: %w", err)

	}

	return user, true, nil
}

func (d *PostgresDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time
	err := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)
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
		`INSERT INTO users(name, phone_number, password)
		 VALUES($1, $2, $3)
		 RETURNING id`,
		u.Name,
		u.PhoneNumber,
		u.HashedPassword,
	).Scan(&u.ID)

	if err != nil {
		return nil, fmt.Errorf("can't execute command: %w", err)
	}

	return &u, nil
}
