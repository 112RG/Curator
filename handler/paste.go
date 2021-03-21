package handler

import (
	"github.com/112RG/Curator/model"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreatePaste(c *gin.Context) {
	if len(c.Request.FormValue("raw")) > 0 {
		id, _ := gonanoid.New(5)
		paste := model.Paste{Id: id, Content: c.Request.FormValue("raw"), Title: c.Request.FormValue("title"), CreatedIp: c.ClientIP(), Owner: c.Request.FormValue("passcode")}
		err := h.PasteService.Create(c, paste)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to create paste ID: %s CONTENT: %s", paste.Id, paste.Content)
			c.AbortWithStatusJSON(400, err)
		} else {
			c.Redirect(302, "/"+id)
		}
	}
}

func (h *Handler) DeletePaste(c *gin.Context) {
	pasteId := c.Param("pId")

	if len(pasteId) > 0 {
		err := h.PasteService.Delete(c, pasteId)

		if err != nil {
			log.Error().Err(err).Msgf("Failed to create paste ID: %s", pasteId)
			c.AbortWithStatusJSON(400, "Failed to delete paste")
		} else {
			c.JSON(200, "Deleted paste")
		}
	}
}

func (h *Handler) TestPaste(c *gin.Context) {
	pasteId := c.Param("pId")

	pastes, err := h.PasteService.GetOwnerPastes(c, pasteId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create paste ID: %s", pasteId)
		c.AbortWithStatusJSON(400, "Failed to delete paste")
	} else {

		c.JSON(200, pastes)
	}
}
