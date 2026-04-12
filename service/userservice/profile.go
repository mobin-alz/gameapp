package userservice

import (
	"github.com/mobin-alz/gameapp/dto"
	"github.com/mobin-alz/gameapp/pkg/richerror"
)

func (s Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.Profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I have not expected the repository call return "record not found" error,
		//because I assume the interactor input is sanitized.
		return dto.ProfileResponse{}, richerror.New(op).WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	// return User
	return dto.ProfileResponse{Name: user.Name}, nil
}
