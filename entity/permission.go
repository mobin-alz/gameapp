package entity

type Permission struct {
	ID    uint
	Title PermissionTitle
}
type PermissionTitle string

const (
	UserListPermission   = PermissionTitle("mysqluser-list")
	UserDeletePermission = PermissionTitle("mysqluser-delete")
)
