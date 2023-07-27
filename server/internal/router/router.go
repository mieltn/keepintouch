package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/handlers/user"
)

func InitRouter(
	user *user.Handler,
) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.GET("/register", user.Register)
	auth.GET("/login", user.Login)
	auth.GET("/logout", user.Logout)

	return r
}