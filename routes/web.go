package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func addWebRoutes(rg *gin.RouterGroup) {
	web := rg.Group("/")
	web.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
	})
	web.GET("/:pId", func(c *gin.Context) {
		pasteId := c.Param("pId")
		if len(pasteId) > 0 && pasteId != "favicon.ico" {
			paste, err := pasteRepository.FindByID(pasteId)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to get paste %s", pasteId)
				c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			} else {
				c.HTML(http.StatusOK, "paste.html", map[string]interface{}{
					"id":      paste.ID,
					"content": paste.Content,
				})
			}
		}
	})
	web.GET("/:pId/raw", func(c *gin.Context) {
		pasteId := c.Param("pId")
		c.Header("Content-Type", "text/plain; charset=utf-8")
		paste, _ := pasteRepository.FindByID(pasteId)
		c.String(200, paste.Content)
	})

}
