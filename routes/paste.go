package routes

import (
	"net/http"

	"github.com/112RG/Curator/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

func addPasteRoutes(rg *gin.RouterGroup) {
	paste := rg.Group("/paste")

	paste.POST("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; charset=utf-8")
		if len(c.Request.FormValue("raw")) > 0 {
			id, _ := gonanoid.New(5)
			paste := &models.Paste{ID: id, Content: c.Request.FormValue("raw")}
			err := pasteRepository.CreatePaste(paste)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to create paste ID: %s CONTENT: %s", paste.ID, paste.Content)
				c.AbortWithStatusJSON(400, err)
			} else {
				c.Redirect(302, "/"+id)
			}
		}

	})
	paste.DELETE("/:pId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
