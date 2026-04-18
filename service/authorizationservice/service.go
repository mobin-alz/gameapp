package authorizationservice

import "github.com/mobin-alz/gameapp/entity"

type Repository interface {
	GetUserPermissionTitles(userID uint) ([]entity.PermissionTitle, error)
}
type Service struct {
}

func (s Service) CheckAccess(userID uint, permissions ...entity.PermissionTitle) (bool, error) {
	// get all ACLs for the given role

	//get all ACLs for the given user

	// merge all check

	// check the access
}
