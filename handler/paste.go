// Package classification Paste API.
//
// A simple clean Paste API
//
// Terms Of Service:
//
//     Schemes: https
//     Host: localhost:5000
//     Version: 0.0.1
//     Contact: Nobdy<nobody@nobody.com>
//
//
// swagger:meta
package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"curator/model"

	"github.com/gorilla/mux"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreatePaste(w http.ResponseWriter, req *http.Request) {
	id, _ := gonanoid.New(5)
	paste := model.Paste{
		Id: id, Content: req.FormValue("raw"),
		Title:       sql.NullString{String: req.FormValue("title")},
		OwnerId:     sql.NullString{String: req.FormValue("passcode")},
		TimeCreated: time.Now()}

	err := h.PasteService.Create(req.Context(), paste)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create paste ID: %s CONTENT: %s", paste.Id, paste.Content)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		log.Debug().Msgf("Sending paste: %s", paste.Id)
		w.Write([]byte(id))
	}
}

func (h *Handler) DeletePaste(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pasteId := vars["pId"]
	if len(pasteId) > 0 {
		err := h.PasteService.Delete(req.Context(), pasteId)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to delete paste ID: %s", pasteId)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Write([]byte("Deleted paste"))
		}
	}
}

func (h *Handler) TestPaste(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ownerId := vars["OId"]

	pastes, err := h.PasteService.GetOwnerPastes(req.Context(), ownerId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get pastes for owner: %s", ownerId)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		js, err := json.Marshal(pastes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
