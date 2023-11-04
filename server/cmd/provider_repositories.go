package main

import (
	repoUser "github.com/mieltn/keepintouch/internal/repositories/mongo/user"
	repoChat "github.com/mieltn/keepintouch/internal/repositories/mongo/chat"
	repoMessage "github.com/mieltn/keepintouch/internal/repositories/mongo/message"
)

func (sp *serviceProvider) GetUserRepository() *repoUser.Repository {
	if sp.repoUser == nil {
		sp.repoUser = repoUser.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoUser
}

func (sp *serviceProvider) GetChatRepository() *repoChat.Repository {
	if sp.repoChat == nil {
		sp.repoChat = repoChat.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoChat
}

func (sp *serviceProvider) GetMessageRepository() *repoMessage.Repository {
	if sp.repoMessage == nil {
		sp.repoMessage = repoMessage.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoMessage
}