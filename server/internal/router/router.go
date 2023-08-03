package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/handlers/chat"
	"github.com/mieltn/keepintouch/internal/handlers/message"
	"github.com/mieltn/keepintouch/internal/handlers/user"
)

func InitRouter(
	user *user.Handler,
	chat *chat.Handler,
	message *message.Handler,
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

	messageGr := r.Group("")
	// messageGr.Use(user.AuthRequired)
	messageGr.GET("/messages/:chatId", message.MessageByChatId)
	messageGr.POST("/messages", message.Create)

	return r
}
