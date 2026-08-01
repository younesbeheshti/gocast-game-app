package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"net/http"
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

func (s *Service) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {

	//TODO: verifying phone number with verification code

	//create new user in storage
	user := entity.User{
		ID:             0,
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
		HashedPassword: getMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	//return created user
	return &dto.RegisterResponse{
		User: dto.UserInfo{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		}}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type LoginResponse struct {
	User   dto.UserInfo `json:"user"`
	Tokens Tokens       `json:"tokens"`
}

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {
	const op = "userservic.Login"
	// TODO - it would be better for user to have two separate methods for existence and getUserBYPhoneNumber

	//check the existence of phone number from repository
	//get the user by phone number
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage("unexpected").WithKind(richerror.KindUnexpected)
	}

	if !exist {
		return nil, fmt.Errorf("username or password is invalid")
	}

	//compare user.pass with req.pass
	if user.HashedPassword != getMD5Hash(req.Password) {
		return nil, fmt.Errorf("username or passwword is invalid")
	}

	// generate jwt
	accessToken, err := s.auth.CreateAccessToken(*user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	refreshToken, err := s.auth.CreateAccessToken(*user)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %w", err)
	}

	return &LoginResponse{User: dto.UserInfo{ID: user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	},
		Tokens: Tokens{AccessToken: accessToken,
			RefreshToken: refreshToken},
	}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])

}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s *Service) GetProfile(req ProfileRequest) (*ProfileResponse, error) {
	const op = "userservice.GetProfile"
	//getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {

		//TODO: can use rich error
		return nil, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"req": req})
	}

	return &ProfileResponse{user.Name}, nil
}

func codeAndMessage(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		return re.Message(), MapKindToHTTPStatusCode(re.Kind())

	default:
		return err.Error(), http.StatusBadRequest
	}

}

func MapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}

}
