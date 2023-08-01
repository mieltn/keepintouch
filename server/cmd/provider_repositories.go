package main

import (
	repoUser "github.com/mieltn/keepintouch/internal/repositories/mongo/user"
)

func (sp *serviceProvider) GetUserRepository() *repoUser.Repository {
	if sp.repoUser == nil {
		sp.repoUser = repoUser.NewRepository(
			sp.GetDB(),
		)
	}
	return sp.repoUser
}
