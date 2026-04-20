package backofficeuserservice

import "github.com/mobin-alz/gameapp/entity"

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) ListAllUsers() ([]entity.User, error) {
	// TODO - implement me
	list := make([]entity.User, 0)

	list = append(list, entity.User{
		ID:          0,
		PhoneNumber: "fake",
		Name:        "fake",
		Password:    "fake",
		Role:        entity.AdminRole,
	})

	return list, nil
}
