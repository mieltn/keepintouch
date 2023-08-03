package main

import (
	"context"
	"log"

	"github.com/mieltn/keepintouch/internal/config"
	hndlUser "github.com/mieltn/keepintouch/internal/handlers/user"
	hndlChat "github.com/mieltn/keepintouch/internal/handlers/chat"
	repoChat "github.com/mieltn/keepintouch/internal/repositories/mongo/chat"
	repoUser "github.com/mieltn/keepintouch/internal/repositories/mongo/user"
	srvChat "github.com/mieltn/keepintouch/internal/services/chat"
	srvJWT "github.com/mieltn/keepintouch/internal/services/jwt"
	srvUser "github.com/mieltn/keepintouch/internal/services/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type OnCloseApp func(context.Context) error

type serviceProvider struct {
	ctx      context.Context
	cfg      *config.Config
	closeFn  []OnCloseApp
	db       *mongo.Client
	repoUser *repoUser.Repository
	repoChat *repoChat.Repository
	srvUser  *srvUser.Service
	srvChat  *srvChat.Service
	srvJWT   *srvJWT.Service
	hndlUser *hndlUser.Handler
	hndlChat *hndlChat.Handler
}

func (sp *serviceProvider) Close() {
	for _, fn := range sp.closeFn {
		if err := fn(sp.ctx); err != nil {
			log.Fatal(err)
		}
	}
}

func (sp *serviceProvider) addError(err error) {
	log.Fatal(err)
}

func (sp *serviceProvider) onClose(fn OnCloseApp) {
	sp.closeFn = append(sp.closeFn, fn)
}

func NewServiceProvider(ctx context.Context, cfg *config.Config) *serviceProvider {
	return &serviceProvider{
		ctx: ctx,
		cfg: cfg,
	}
}
