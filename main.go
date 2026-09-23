package main

import (
	"context"
	"fmt"
	presenceClient "github.com/younesbeheshti/gocast_game/adapter/presence"
	"github.com/younesbeheshti/gocast_game/adapter/redis"
	"github.com/younesbeheshti/gocast_game/config"
	"github.com/younesbeheshti/gocast_game/delivery/httpserver"
	"github.com/younesbeheshti/gocast_game/logger"
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
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "net/http/pprof"
)

func main() {

	go func() {
		// TODO: add enabler config variable
		// curl http://localhost:8099/debug/pprof/goroutine --output goroutings.o
		// go tool pprof -http=:8081 goroutings.o
		http.ListenAndServe(":8099", nil)
	}()

	//TODO: read cofig path from command line

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger.Named("main").Info("config", zap.Any("config", cfg))

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

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		sch := scheduler.New(matchingSvc, cfg.Scheduler)
		sch.Start(ctx, &wg)
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

	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.PresenceService, presenceRepo)

	// TODO: panic - replace presenceSve with presence grpc client

	presenceAdapter, _ := presenceClient.New(":8000")

	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo, presenceAdapter, redisAdapter)

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
