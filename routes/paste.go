package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPasteRoutes(rg *gin.RouterGroup) {
	paste := rg.Group("/paste")

	paste.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	paste.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	paste.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
