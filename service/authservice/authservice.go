package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/younesbeheshti/gocast_game/entity"
	"strings"
	"time"
)

type Service struct {
	signKey            string
	accessDurationTime time.Duration
	refreshDuration    time.Duration
	accessSubject      string
	refreshSubject     string
}

func New(signKey, accessSubject, refreshSubject string, accessDurationTime, refreshDurationTime time.Duration) *Service {
	return &Service{
		signKey:            signKey,
		accessDurationTime: accessDurationTime,
		refreshDuration:    refreshDurationTime,
		accessSubject:      accessSubject,
		refreshSubject:     refreshSubject,
	}
}

func (s *Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.accessSubject, s.accessDurationTime)
}
func (s *Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.refreshSubject, s.refreshDuration)
}
func (s *Service) ParseToken(tokenString string) (*Claims, error) {

	tokenStr := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
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

func (s *Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}

	return t.SignedString([]byte(s.signKey))
}
