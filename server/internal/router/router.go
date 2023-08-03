package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/handlers/user"
	"github.com/mieltn/keepintouch/internal/handlers/chat"
)

func InitRouter(
	user *user.Handler,
	chat *chat.Handler,
) *gin.Engine {
	r := gin.Default()

	authGr := r.Group("/auth")
	authGr.POST("/register", user.Register)
	authGr.POST("/login", user.Login)
	authGr.POST("/logout", user.Logout)
	authGr.POST("/refresh", user.Refresh)

	chatGr := r.Group("")
	// chatGr.Use(user.AuthRequired)
	chatGr.GET("/chats", chat.List)
	chatGr.POST("/chats", chat.Create)

	return r
}