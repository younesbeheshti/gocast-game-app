package postgres

import (
	"database/sql"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"time"
)

func (d PostgresDB) GetUserByID(userID uint) (*entity.User, error) {
	row := d.db.QueryRow(`select * from users where id=$1`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("cant scan query result: %w", err)

	}

	return user, nil
}

func (d PostgresDB) GetUserByPhoneNumber(phoneNumber string) (*entity.User, bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber)
	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("cant scan query result: %w", err)

	}

	return user, true, nil
}

func (d PostgresDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber)
	_, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("cant scan query result: %w", err)
	}
	return false, nil
}
func (d PostgresDB) Register(u entity.User) (*entity.User, error) {
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
func scanUser(rows *sql.Row) (*entity.User, error) {
	var user entity.User
	var createdAt time.Time
	err := rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.HashedPassword)
	return &user, err
}
