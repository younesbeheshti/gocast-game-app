package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"io"
	"net/http"
	"time"
)

const (
	JwtSecret                  = "secret"
	AccessToken                = "access_token"
	RefreshToken               = "refresh_token"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	userId := r.Context().Value("user_id")

	req := userservice.ProfileRequest{UserID: userId.(uint)}

	psqlRepo := postgres.New()
	authSvc := authservice.New(JwtSecret, AccessToken, RefreshToken, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	userSvc := userservice.New(psqlRepo, authSvc)

	resp, err := userSvc.GetProfile(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err.Error())))
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err.Error())))
		return
	}

	w.Write(data)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Only POST method is allowed")
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var req userservice.LoginRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	psqlRepo := postgres.New()
	authSvc := authservice.New(JwtSecret, AccessToken, RefreshToken, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	userSvc := userservice.New(psqlRepo, authSvc)
	resp, err := userSvc.Login(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	data, err = json.Marshal(resp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	w.Write(data)
}

func main() {

	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:            JwtSecret,
			AccessSubject:      AccessToken,
			RefreshSubject:     RefreshToken,
			AccessDurationTime: AccessTokenExpireDuration,
			RefreshDuration:    RefreshTokenExpireDuration,
		},
		Psql: postgres.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "postgres",
			Sslmode:  "disable",
		}}

	userSvc, authSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()

	http.HandleFunc("/users/login", userLoginHandler)
	http.Handle(
		"/users/profile",
		Middleware(http.HandlerFunc(userProfileHandler)),
	)

}

func setupServices(cfg config.Config) (*userservice.Service, *authservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	psqlRepo := postgres.New(cfg.Psql)
	userSvc := userservice.New(psqlRepo, authSvc)

	return userSvc, authSvc
}

func testDatabase() {
	psql := postgres.New()
	u, err := psql.Register(entity.User{ID: 0, Name: "ali", PhoneNumber: "09123223"})
	fmt.Println(u, err)

	isUnique, err := psql.IsPhoneNumberUnique(u.PhoneNumber + "23")
	fmt.Println(isUnique, err)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		authSvc := authservice.New(JwtSecret, AccessToken, RefreshToken, AccessTokenExpireDuration, RefreshTokenExpireDuration)

		claims, err := authSvc.ParseToken(authToken)
		fmt.Println(claims, err)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized")
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"user_id",
			claims.UserID,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
