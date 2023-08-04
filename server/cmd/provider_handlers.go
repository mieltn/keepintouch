package main

import (
	hndlChat "github.com/mieltn/keepintouch/internal/handlers/chat"
	hndlMessage "github.com/mieltn/keepintouch/internal/handlers/message"
	hndlUser "github.com/mieltn/keepintouch/internal/handlers/user"
	hndlWs "github.com/mieltn/keepintouch/internal/handlers/ws"
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

func (sp *serviceProvider) GetMessageHandler() *hndlMessage.Handler {
	if sp.hndlMessage == nil {
		sp.hndlMessage = hndlMessage.NewHandler(
			sp.GetMessageRepository(),
		)
	}
	return sp.hndlMessage
}

func (sp *serviceProvider) GetWsHandler() *hndlWs.Handler {
	if sp.hndlWs == nil {
		sp.hndlWs = hndlWs.NewHandler(
			sp.GetUserRepository(),
			sp.GetMessageRepository(),
			sp.GetBroadcasterService(),
		)
	}
	return sp.hndlWs
}
