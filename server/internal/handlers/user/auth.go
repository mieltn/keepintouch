package user

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mieltn/keepintouch/internal/dto"
)

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
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errNoTokenInHeader.Error()})
		c.Abort()
		return
	}
	token := strings.Split(tokenHeader, " ")
	if isValid, err := h.service.Validate(c, token[len(token)-1]); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

func (h *Handler) Logout(c *gin.Context) {}