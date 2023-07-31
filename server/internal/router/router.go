package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/handlers/user"
)

func InitRouter(
	user *user.Handler,
) *gin.Engine {
	r := gin.Default()

	authGr := r.Group("/auth")
	authGr.Use(user.AuthRequired)
	authGr.POST("/register", user.Register)
	authGr.POST("/login", user.Login)
	authGr.POST("/logout", user.Logout)
	authGr.POST("/refresh", user.Refresh)

	return r
}