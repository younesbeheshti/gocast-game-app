package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/younesbeheshti/gocast_game/entity"
)

type Repository interface {
	Register(u entity.User) (*entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*entity.User, bool, error)
	GetUserByID(uint) (*entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(u entity.User) (string, error)
	CreateRefreshToken(u entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(repo Repository, authGenerator AuthGenerator) Service {
	return Service{repo: repo, auth: authGenerator}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])

}
