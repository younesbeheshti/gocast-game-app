package main

import (
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"time"
)

const (
	JwtSecret                  = "secret"
	AccessToken                = "access_token"
	RefreshToken               = "refresh_token"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

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

	// TODO - add command for apply
	//mgr := migrator.New(cfg.Psql)
	//mgr.Up()

	userSvc, authSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()

}

func setupServices(cfg config.Config) (userservice.Service, authservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	psqlRepo := postgres.New(cfg.Psql)
	userSvc := userservice.New(psqlRepo, authSvc)

	return userSvc, authSvc
}

//func testDatabase() {
//	psql := postgres.New()
//	u, err := psql.Register(entity.User{ID: 0, Name: "ali", PhoneNumber: "09123223"})
//	fmt.Println(u, err)
//
//	isUnique, err := psql.IsPhoneNumberUnique(u.PhoneNumber + "23")
//	fmt.Println(isUnique, err)
//}

//func Middleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authToken := r.Header.Get("Authorization")
//		authSvc := authservice.New(JwtSecret, AccessToken, RefreshToken, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//		claims, err := authSvc.ParseToken(authToken)
//		fmt.Println(claims, err)
//		if err != nil {
//			w.WriteHeader(http.StatusUnauthorized)
//			fmt.Fprint(w, "Unauthorized")
//			return
//		}
//
//		ctx := context.WithValue(
//			r.Context(),
//			"user_id",
//			claims.UserID,
//		)
//
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
