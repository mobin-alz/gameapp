package userservice

import (
	"github.com/mobin-alz/gameapp/param"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I have not expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized.
		return param.ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	// return User
	return param.ProfileResponse{Name: user.Name}, nil
}
