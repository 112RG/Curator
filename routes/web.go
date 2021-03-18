package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addWebRoutes(rg *gin.RouterGroup) {
	web := rg.Group("/")
	web.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"id":      "tAAAAAAAAAAAAest",
			"content": "taaaaaaaaaaest",
		})
	})
	web.GET("/:pId", func(c *gin.Context) {
		pasteId := c.Param("pId")
		if len(pasteId) > 0 && pasteId != "favicon.ico" {
			pasteId := c.Param("pId")
			paste, _ := pasteRepository.FindByID(pasteId)
			c.HTML(http.StatusOK, "paste.html", map[string]interface{}{
				"id":      paste.ID,
				"content": paste.Content,
			})
		}
	})
	web.GET("/:pId/raw", func(c *gin.Context) {
		pasteId := c.Param("pId")
		c.Header("Content-Type", "text/plain; charset=utf-8")
		paste, _ := pasteRepository.FindByID(pasteId)
		c.String(200, paste.Content)
	})

}
