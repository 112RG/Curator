package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "test",
	})
}
