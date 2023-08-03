package main

import (
	"context"
	"log"
	"time"

	"github.com/mieltn/keepintouch/internal/config"
	"github.com/mieltn/keepintouch/internal/router"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	sp := NewServiceProvider(ctx, cfg)
	defer sp.Close()

	r := router.InitRouter(sp.GetUserHandler(), sp.GetChatHandler())
	r.Run()
}