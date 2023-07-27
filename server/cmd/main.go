package main

import (
	"context"
	"log"
	"time"

	"github.com/mieltn/keepintouch/db"
	"github.com/mieltn/keepintouch/internal/router"
	userRepo "github.com/mieltn/keepintouch/internal/repositories/mongo/user"
	userSrv "github.com/mieltn/keepintouch/internal/services/user"
	userHandler "github.com/mieltn/keepintouch/internal/handlers/user"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	db, err := db.NewMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := userRepo.NewRepository(db.GetDB())
	userSrv := userSrv.NewService(userRepo)
	userHandler := userHandler.NewHandler(userSrv)

	r := router.InitRouter(userHandler)
	r.Run()
}