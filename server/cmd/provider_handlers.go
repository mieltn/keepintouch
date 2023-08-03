package main

import (
	hndlUser "github.com/mieltn/keepintouch/internal/handlers/user"
	hndlChat "github.com/mieltn/keepintouch/internal/handlers/chat"
)

func (sp *serviceProvider) GetUserHandler() *hndlUser.Handler {
	if sp.hndlUser == nil {
		sp.hndlUser = hndlUser.NewHandler(
			sp.GetUserService(),
		)
	}
	return sp.hndlUser
}

func (sp *serviceProvider) GetChatHandler() *hndlChat.Handler {
	if sp.hndlChat == nil {
		sp.hndlChat = hndlChat.NewHandler(
			sp.GetChatService(),
		)
	}
	return sp.hndlChat
}
