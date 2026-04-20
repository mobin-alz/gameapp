package authorizationservice

import (
	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

type Repository interface {
	GetUserPermissionTitles(userID uint, role entity.Role) ([]entity.PermissionTitle, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CheckAccess(userID uint, role entity.Role, permissions ...entity.PermissionTitle) (bool, error) {
	const op = "authorizationservice.CheckAccess"
	permissionTitles, err := s.repo.GetUserPermissionTitles(userID, role)

	if err != nil {
		return false, richerror.New(op).WithError(err)
	}

	for _, p := range permissions {
		for _, pt := range permissionTitles {
			if pt == p {
				return true, nil
			}
		}
	}

	// check the access
	return false, nil
}
