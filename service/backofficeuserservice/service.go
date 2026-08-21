package backofficeuserservice

import "github.com/younesbeheshti/gocast_game/entity"

type Service struct {
}

func New() Service {
	return Service{}
}

func (s *Service) DoSomething() ([]entity.User, error) {

	list := make([]entity.User, 0)

	list = append(list, entity.User{
		ID:             1,
		PhoneNumber:    "fake",
		Name:           "fake",
		HashedPassword: "fake",
		Role:           entity.AdminRole,
	})

	return list, nil

}
