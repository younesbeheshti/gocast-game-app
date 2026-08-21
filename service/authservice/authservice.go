package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/younesbeheshti/gocast_game/entity"
	"strings"
	"time"
)

type Config struct {
	SignKey            string        `koanf:"sign_key"`
	AccessDurationTime time.Duration `koanf:"access_duration_time"`
	RefreshDuration    time.Duration `koanf:"refresh_duration"`
	AccessSubject      string        `koanf:"access_subject"`
	RefreshSubject     string        `koanf:"refresh_subject"`
}

type Service struct {
	config Config
}

func New(config Config) Service {
	return Service{
		config: config,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.AccessSubject, s.config.AccessDurationTime)
}
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.RefreshSubject, s.config.RefreshDuration)
}
func (s Service) ParseToken(tokenString string) (*Claims, error) {

	tokenStr := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (s Service) createToken(userID uint, role entity.Role, subject string, expireDuration time.Duration) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
		Role:   role,
	}

	return t.SignedString([]byte(s.config.SignKey))
}
