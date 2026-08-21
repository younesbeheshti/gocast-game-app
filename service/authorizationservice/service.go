package authorizationservice

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
)

type Repository interface {
	GetUserPermissionTitles(userID uint, role entity.Role) ([]entity.PermissionTitle, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckAccess(userID uint, role entity.Role, permissions ...entity.PermissionTitle) (bool, error) {

	const op = "authorizationservice.CheckAccess"

	permissionTitles, err := s.repo.GetUserPermissionTitles(userID, role)
	if err != nil {
		fmt.Println("err : ", err.Error())
		return false, richerror.New(op).WithErr(err)
	}

	// check the access
	for _, pt := range permissionTitles {
		for _, p := range permissions {
			if p == pt {
				return true, nil
			}
		}
	}

	return false, nil
}
