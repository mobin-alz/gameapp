package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.db.QueryRow("SELECT * FROM users WHERE phone_number = ?", phoneNumber)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, richerror.New(op).
			WithError(err).
			WithMessage(errmsg.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}
	return false, nil
}

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number,password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.db.QueryRow("SELECT * FROM users WHERE phone_number = ?", phoneNumber)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).
				WithError(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}
		// TODO- log unexpected error for better observability
		return entity.User{}, richerror.New(op).
			WithError(err).
			WithMessage(errmsg.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func (d *MySQLDB) GetUserByID(userID uint) (entity.User, error) {

	row := d.db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	const op = "mysql.GetUserByID"

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).WithError(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New(op).
			WithError(err).
			WithMessage(errmsg.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func scanUser(scanner Scanner) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password)

	return user, err
}
