package main

import (
	"context"
	"fmt"
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver"
	"github.com/younesbeheshti/gocast_game/repository/migrator"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	psqlaccesscontrol "github.com/younesbeheshti/gocast_game/repository/postgres/accesscontrol"
	"github.com/younesbeheshti/gocast_game/repository/postgres/user"
	"github.com/younesbeheshti/gocast_game/repository/redis/redismatching"
	"github.com/younesbeheshti/gocast_game/repository/redis/redispresence"
	"github.com/younesbeheshti/gocast_game/scheduler"
	"github.com/younesbeheshti/gocast_game/service/authorizationservice"
	"github.com/younesbeheshti/gocast_game/service/authservice"
	"github.com/younesbeheshti/gocast_game/service/backofficeuserservice"
	"github.com/younesbeheshti/gocast_game/service/matchingservice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"github.com/younesbeheshti/gocast_game/validator/matchingvalidator"
	"github.com/younesbeheshti/gocast_game/validator/uservalidator"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//TODO: read cofig path from command line

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	// TODO - add command for apply
	mgr := migrator.New(cfg.Psql)
	mgr.Up()

	// TODO - create struct and add these returned items as struct field
	userSvc, authSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV, presenceSvc := setupServices(cfg)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		sch := scheduler.New(matchingSvc, cfg.Scheduler)
		sch.Start(ctx)
	}()

	server := httpserver.
		New(cfg, authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV, presenceSvc)

	if err := server.Serve(ctx); err != nil {
		log.Printf("HTTP server stopped: %v", err)
	}

	//sigchnl := make(chan os.Signal, 1)
	//signal.Notify(sigchnl, syscall.SIGINT)
	//<-sigchnl
	fmt.Println("Shutting down...")

}

func setupServices(cfg config.Config) (
	userservice.Service, authservice.Service, uservalidator.Validator,
	backofficeuserservice.Service, authorizationservice.Service,
	matchingservice.Service, matchingvalidator.Validator, presenceservice.Service) {

	authSvc := authservice.New(cfg.Auth)
	psqlRepo := postgres.New(cfg.Psql)

	userPsql := psqluser.New(&psqlRepo)

	userSvc := userservice.New(userPsql, authSvc)

	uV := uservalidator.New(userPsql)
	backofficeUserSvc := backofficeuserservice.New()

	aclPsql := psqlaccesscontrol.New(&psqlRepo)
	authorizationSvc := authorizationservice.New(aclPsql)

	matchingV := matchingvalidator.New()

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)
	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.PresenceService, presenceRepo)

	return userSvc, authSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV, presenceSvc
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
