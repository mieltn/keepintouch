package message

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/dto"
)

var (
	errBadChatId = errors.New("wrong chat id query parameter")
)

type MessageRepository interface {
	MessageByChatId(context.Context, *dto.MessageByChatIdReq) ([]*dto.Message, error)
	Create(context.Context, *dto.MessageCreateReq) (*dto.Message, error)
}

type Handler struct {
	repo MessageRepository
}

func NewHandler(repo MessageRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) MessageByChatId(c *gin.Context) () {
	ctx, cancel := context.WithTimeout(c, time.Second * 10)
	defer cancel()

	var req dto.MessageByChatIdReq
	params := c.Request.URL.Query()
	
	req.Id = c.Param("chatId")
	log.Println(req.Id)
	
	if limit, ok := params["limit"]; !ok {
		req.Limit = 25
	} else {
		limitInt, err := strconv.Atoi(limit[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Limit = limitInt
	}

	if offset, ok := params["offset"]; !ok {
		req.Offset = 0
	} else {
		offsetInt, err := strconv.Atoi(offset[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Offset = offsetInt
	}

	messages, err := h.repo.MessageByChatId(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *Handler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second * 10)
	defer cancel()

	var req dto.MessageCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat, err := h.repo.Create(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chat)
}
