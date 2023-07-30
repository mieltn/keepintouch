package main

import (
	"context"
	"log"
	"time"

	"github.com/mieltn/keepintouch/internal/config"
	"github.com/mieltn/keepintouch/db"
	userHandler "github.com/mieltn/keepintouch/internal/handlers/user"
	userRepo "github.com/mieltn/keepintouch/internal/repositories/mongo/user"
	"github.com/mieltn/keepintouch/internal/router"
	userSrv "github.com/mieltn/keepintouch/internal/services/user"
	JWTSrv "github.com/mieltn/keepintouch/internal/services/jwt"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	JWTService := JWTSrv.NewService(cfg)

	userRepo := userRepo.NewRepository(db.GetDB())
	userSrv := userSrv.NewService(userRepo, JWTService)
	userHandler := userHandler.NewHandler(userSrv)

	r := router.InitRouter(userHandler)
	r.Run()
}