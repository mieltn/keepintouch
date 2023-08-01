package main

import (
	hndlUser "github.com/mieltn/keepintouch/internal/handlers/user"
)

func (sp *serviceProvider) GetUserHandler() *hndlUser.Handler {
	if sp.hndlUser == nil {
		sp.hndlUser = hndlUser.NewHandler(
			sp.GetUserService(),
		)
	}
	return sp.hndlUser
}
