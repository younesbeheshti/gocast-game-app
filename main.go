package main

import (
	"context"
	"encoding/json"
	"fmt"
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
func userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
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

	var req userservice.RegisterRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	psqlRepo := postgres.New()
	authSvc := authservice.New(JwtSecret, AccessToken, RefreshToken, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	userSvc := userservice.New(psqlRepo, authSvc)

	fmt.Println(req)
	_, err = userSvc.Register(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"success": "%s"}`, "OK")))

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
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.Handle(
		"/users/profile",
		middleware(http.HandlerFunc(userProfileHandler)),
	)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func testDatabase() {
	psql := postgres.New()
	u, err := psql.Register(entity.User{ID: 0, Name: "ali", PhoneNumber: "09123223"})
	fmt.Println(u, err)

	isUnique, err := psql.IsPhoneNumberUnique(u.PhoneNumber + "23")
	fmt.Println(isUnique, err)
}

func middleware(next http.Handler) http.Handler {
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
