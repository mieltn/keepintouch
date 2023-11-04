package chat

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/dto"
)

type ChatService interface {
	List(context.Context, *dto.ChatListReq) ([]*dto.Chat, error)
	Create(context.Context, *dto.ChatCreateReq) (*dto.Chat, error)
}

type Handler struct {
	service ChatService
}

func NewHandler(service ChatService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(c *gin.Context) () {

	var req dto.ChatListReq
	params := c.Request.URL.Query()
	
	if ids, ok := params["ids"]; !ok {
		req.Ids = []string{}
	} else {
		req.Ids = ids
	}
	
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

	chats, err := h.service.List(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (h *Handler) Create(c *gin.Context) {

	var req dto.ChatCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat, err := h.service.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chat)
}
