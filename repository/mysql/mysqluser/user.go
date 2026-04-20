package mysqluser

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"github.com/mobin-alz/gameapp/repository/mysql"
)

func (d DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.conn.Conn().QueryRow("SELECT * FROM users WHERE phone_number = ?", phoneNumber)

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

func (d DB) Register(u entity.User) (entity.User, error) {
	res, err := d.conn.Conn().Exec(`insert into users(name, phone_number,password,role) values(?, ?, ?, ?)`,
		u.Name, u.PhoneNumber, u.Password, u.Role.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}

func (d DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.conn.Conn().QueryRow("SELECT * FROM users WHERE phone_number = ?", phoneNumber)

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

func (d DB) GetUserByID(userID uint) (entity.User, error) {

	row := d.conn.Conn().QueryRow("SELECT * FROM users WHERE id = ?", userID)
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

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	var roleStr string
	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)

	user.Role = entity.MapToRoleEntity(roleStr)

	return user, err
}
