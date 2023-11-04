package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mieltn/keepintouch/internal/dto"
)

type UserRepository interface {
	GetUserById(context.Context, string) (*dto.User, error)
}

type MessageRepository interface {
	Create(context.Context, *dto.MessageCreateReq) (*dto.Message, error)
}

type BroadcasterService interface {
	Run()
	GetBroadcast() chan *dto.Message
	GetRegister() chan *dto.Client
	GetUnregister() chan *dto.Client
}

type Handler struct {
	userRepo UserRepository
	msgRepo MessageRepository
	broadcasterSrv BroadcasterService
}

func NewHandler(userRepo UserRepository, msgRepo MessageRepository, broadcasterSrv BroadcasterService) *Handler {
	return &Handler{
		userRepo: userRepo,
		msgRepo: msgRepo,
		broadcasterSrv: broadcasterSrv,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinChat(c *gin.Context) {

	var req dto.JoinChatReq
	req.ChatId = c.Query("chatId")
	req.UserId = c.Query("userId")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &dto.Client{
		ChatId: req.ChatId,
		UserId: req.UserId,
		Conn: conn,
	}
	h.broadcasterSrv.GetRegister() <- client

	user, err := h.userRepo.GetUserById(c, req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	msg, err := h.msgRepo.Create(c, &dto.MessageCreateReq{
		ChatId: req.ChatId,
		UserId: req.UserId,
		Text: fmt.Sprintf("%s joined the chat", user.Username),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.broadcasterSrv.GetBroadcast() <- msg
	if err = h.receive(c, client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) receive(ctx context.Context, client *dto.Client) error {
	defer func() {
		h.broadcasterSrv.GetUnregister() <- client
		client.Conn.Close()
	}()

	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			return err
		}

		var msg dto.MessageCreateReq
		if err := json.Unmarshal(p, &msg); err != nil {
			return err
		}

		message, err := h.msgRepo.Create(ctx, &msg)
		if err != nil {
			return err
		}

		h.broadcasterSrv.GetBroadcast() <- message
	}
}
