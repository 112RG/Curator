package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addWebRoutes(rg *gin.RouterGroup) {
	web := rg.Group("/")

	web.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"now": "test",
		})
	})
}
