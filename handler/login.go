package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetLogin(w http.ResponseWriter, req *http.Request) {
	h.TemplateService.ExecuteTemplate(w, "login.html", map[string]interface{}{})
}
func (h *Handler) PostLogin(w http.ResponseWriter, req *http.Request) {
	session, err := h.SessionService.Get(req, "session")
	isLoggedIn := session.Values["loggedin"]

	if isLoggedIn != "true" {
		log.Debug().Msg("Not logged in")
		if err != nil {
			log.Error().Err(err).Msg("Failed to get paste")
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Debug().Msg(req.FormValue("username"))
		session.Values["loggedin"] = "true"
		session.Values["username"] = req.FormValue("username")
		session.Save(req, w)
		http.Redirect(w, req, "/", 302)
	} else {
		log.Debug().Msg("Logged in")
		http.Redirect(w, req, "/", 302)
	}

}

func (h *Handler) checkLogin(w http.ResponseWriter, req *http.Request) (bool, *sessions.Session) {
	session, err := h.SessionService.Get(req, "session")
	isLoggedIn := session.Values["loggedin"]
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return false, nil
	}
	if isLoggedIn != "true" {
		return false, nil
	} else {
		log.Debug().Msg("logged in")
		return true, session
	}

}
