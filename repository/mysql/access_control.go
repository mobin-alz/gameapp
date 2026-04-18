package mysql

import (
	"database/sql"
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"github.com/mobin-alz/gameapp/pkg/slice"
	"strings"
	"time"
)

func (d *MySQLDB) GetUserPermissionTitles(userID uint) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionTitles"
	user, err := d.GetUserByID(userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	roleACL := make([]entity.AccessControl, 0)
	rows, err := d.db.Query("SELECT * FROM access_controls WHERE actor_type= ? AND actor_id=?", entity.RoleActorType, user.Role)
	if err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = richerror.New(op).WithError(err).WithMessage("failed to close rows")
		}
	}(rows)

	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		roleACL = append(roleACL, acl)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	userACl := make([]entity.AccessControl, 0)
	userRows, err := d.db.Query("SELECT * FROM access_controls WHERE actor_type= ? AND actor_id=?", entity.UserActorType, user.ID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer func(userRows *sql.Rows) {
		err := userRows.Close()
		if err != nil {
			err = richerror.New(op).WithError(err).WithMessage("failed to close rows")
		}
	}(userRows)

	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		userACl = append(userACl, acl)
	}

	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	// merge ACLs by permission id
	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	for _, u := range userACl {
		if !slice.DoesExist(permissionIDs, u.PermissionID) {
			permissionIDs = append(permissionIDs, u.PermissionID)
		}
	}

	args := make([]interface{}, len(permissionIDs))
	for i, id := range permissionIDs {
		args[i] = id
	}
	if len(permissionIDs) == 0 {
		return nil, nil
	}

	query := "SELECT * FROM permissions WHERE id IN (?" +
		strings.Repeat(",?", len(permissionIDs)-1) + ")"
	rows, err = d.db.Query(query, args...)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = richerror.New(op).WithError(err).WithMessage("failed to close rows")
		}
	}(rows)

	if err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	permissionTitles := make([]entity.PermissionTitle, 0)

	for rows.Next() {
		permission, err := scanPermission(rows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return permissionTitles, nil

}

func scanAccessControl(scanner Scanner) (entity.AccessControl, error) {
	var createdAt time.Time
	var acl entity.AccessControl

	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
	return acl, err
}

func scanPermission(scanner Scanner) (entity.Permission, error) {
	var createdAt time.Time
	var permission entity.Permission

	err := scanner.Scan(&permission.ID, &permission.Title, &createdAt)
	return permission, err
}
