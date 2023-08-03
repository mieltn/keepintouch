package main

import (
	srvJWT "github.com/mieltn/keepintouch/internal/services/jwt"
	srvUser "github.com/mieltn/keepintouch/internal/services/user"
	srvChat "github.com/mieltn/keepintouch/internal/services/chat"
)

func (sp *serviceProvider) GetJWTService() *srvJWT.Service {
	if sp.srvJWT == nil {
		sp.srvJWT = srvJWT.NewService(
			sp.cfg,
		)
	}
	return sp.srvJWT
}

func (sp *serviceProvider) GetUserService() *srvUser.Service {
	if sp.srvUser == nil {
		sp.srvUser = srvUser.NewService(
			sp.GetUserRepository(),
			sp.GetJWTService(),
		)
	}
	return sp.srvUser
}

func (sp *serviceProvider) GetChatService() *srvChat.Service {
	if sp.srvChat == nil {
		sp.srvChat = srvChat.NewService(
			sp.GetChatRepository(),
		)
	}
	return sp.srvChat
}
