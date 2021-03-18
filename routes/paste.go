package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/112RG/Curator/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func addPasteRoutes(rg *gin.RouterGroup) {
	paste := rg.Group("/paste")

	paste.POST("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; charset=utf-8")
		id, _ := gonanoid.New(5)

		if len(c.Request.FormValue("raw")) > 0 {
			paste := &models.Paste{ID: id, Content: c.Request.FormValue("raw")}
			fmt.Println(paste.ID)
			fmt.Println(paste.Content)
			err := pasteRepository.CreatePaste(paste)
			if err != nil {
				log.Fatalln(err.Error())
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
