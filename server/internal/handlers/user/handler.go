package user

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(context.Context)
	Login(context.Context)
	Logout(context.Context)
}

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(*gin.Context) {}
func (h *Handler) Login(*gin.Context) {}
func (h *Handler) Logout(*gin.Context) {}