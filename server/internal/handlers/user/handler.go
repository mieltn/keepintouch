package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/dto"
)

var (
	errNoTokenInHeader = errors.New("no token found in header")
)

type UserService interface {
	Register(context.Context, *dto.UserCreateReq) (*dto.UserCreateRes, error)
	Login(context.Context, *dto.UserLoginReq) (*dto.UserAuthRes, error)
	Refresh(context.Context, *dto.UserRefreshReq) (*dto.UserAuthRes, error)
	Validate(context.Context, string) (bool, error)
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

func (h *Handler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second * 10)
	defer cancel()

	var req dto.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second * 10)
	defer cancel()

	var req dto.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) Refresh(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second * 10)
	defer cancel()

	var req dto.UserRefreshReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Refresh(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)	
}

func (h *Handler) AuthRequired(c *gin.Context) {
	tokenParam := c.Request.Header["Token"]
	if len(tokenParam) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errNoTokenInHeader.Error()})
		c.Abort()
		return
	}
	if isValid, err := h.service.Validate(c, tokenParam[0]); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

func (h *Handler) Logout(c *gin.Context) {}