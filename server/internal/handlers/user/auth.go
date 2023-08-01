package user

import (
	"context"
	"net/http"
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