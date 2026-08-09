package postgres

import (
	"database/sql"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"time"
)

func (d PostgresDB) GetUserByID(userID uint) (*entity.User, error) {
	const op = "psql.GetUserByID"
	row := d.db.QueryRow(`select * from users where id=$1`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).WithMessage("record not found").WithKind(richerror.KindNotFound)

		}
		return nil, richerror.New(op).WithErr(err).WithMessage("can't scan query result").WithKind(richerror.KindUnexpected)

	}

	return user, nil
}

func (d PostgresDB) GetUserByPhoneNumber(phoneNumber string) (*entity.User, error) {
	const op = "psql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber)
	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).WithMessage("record not found").WithKind(richerror.KindNotFound)
		}

		// TODO: log unexpected error for better observability
		return nil, richerror.New(op).WithErr(err).WithMessage("can't scan the query result").WithKind(richerror.KindUnexpected)

	}

	return user, nil
}

func (d PostgresDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	const op = "psql.IsPhoneNumberUnique"

	row := d.db.QueryRow(`select * from users where phone_number=$1`, phoneNumber)
	_, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage("can't scan the query result").WithKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (d PostgresDB) Register(u entity.User) (*entity.User, error) {
	const op = "psql.Register"

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
