package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPaste(c *gin.Context) {
	pasteId := c.Param("pId")

	if len(pasteId) > 0 && pasteId != "favicon.ico" {
		paste, err := h.PasteService.Get(c, pasteId)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to get paste %s", pasteId)
			c.HTML(http.StatusNotFound, "error.html", map[string]interface{}{
				"message": "Unable to find paste",
			})
		} else {
			c.HTML(http.StatusOK, "paste.html", map[string]interface{}{
				"id":      paste.Id,
				"content": paste.Content,
			})
		}
	}
}

func (h *Handler) GetPasteRaw(c *gin.Context) {
	pasteId := c.Param("pId")
	paste, err := h.PasteService.Get(c, pasteId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get paste %s", pasteId)
		c.HTML(http.StatusNotFound, "error.html", map[string]interface{}{
			"message": "Unable to find paste",
		})
	} else {
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.String(200, paste.Content)
	}
}

func (h *Handler) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
}
