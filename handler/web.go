package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPaste(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	pasteId := vars["pId"]

	if len(pasteId) > 0 && pasteId != "favicon.ico" {
		paste, err := h.PasteService.Get(req.Context(), pasteId)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get paste")
			w.WriteHeader(http.StatusNotFound)
			h.TemplateService.ExecuteTemplate(w, "error.html", map[string]interface{}{})
		} else {
			log.Debug().Msg(paste.OwnerId.String)
			log.Debug().Msg(paste.Title.String)
			var candelete = "cant"
			isLoggedin, session := h.checkLogin(w, req)
			if isLoggedin {
				if session.Values["username"].(string) == paste.OwnerId.String {
					candelete = "can"
				}
			}
			h.TemplateService.ExecuteTemplate(w, "paste.html", map[string]interface{}{
				"id":        paste.Id,
				"content":   paste.Content,
				"owner":     paste.OwnerId.String,
				"date":      paste.TimeCreated,
				"candelete": candelete,
				"title":     paste.Title.String,
				"lang":      paste.Lang,
			})
		}
	}
}

func (h *Handler) GetPasteRaw(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	paste, err := h.PasteService.Get(req.Context(), vars["pId"])
	if err != nil {
		log.Error().Err(err).Msg("Failed to get paste")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(paste.Content))
}
func (h *Handler) GetIndex(w http.ResponseWriter, req *http.Request) {
	isLoggedIn, session := h.checkLogin(w, req)
	if isLoggedIn {
		h.TemplateService.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"username": session.Values["username"],
		})
	} else {
		h.TemplateService.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"username": "",
		})
	}
}
