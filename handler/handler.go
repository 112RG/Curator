package handler

import (
	"net/http"
	"text/template"
	"time"

	"curator/model"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gorilla/mux"
	"github.com/michaeljs1990/sqlitestore"
)

// Handler struct holds required services for handler to function
type Handler struct {
	PasteService    model.PasteService
	TemplateService *template.Template
	SessionService  *sqlitestore.SqliteStore
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *mux.Router
	PasteService    model.PasteService
	TemplateService *template.Template
	SessionService  *sqlitestore.SqliteStore
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {

	// Create a handler (which will later have injected services)
	h := &Handler{
		PasteService:    c.PasteService,
		TemplateService: c.TemplateService,
		SessionService:  c.SessionService,
	}
	c.R.HandleFunc("/login", h.GetLogin).Methods("GET")
	c.R.HandleFunc("/login", h.PostLogin).Methods("POST")

	c.R.HandleFunc("/", h.GetIndex).Methods("GET")

	c.R.HandleFunc("/{pId}", h.GetPaste).Methods("GET")
	c.R.HandleFunc("/{pId}/raw", h.GetPasteRaw).Methods("GET")
	c.R.HandleFunc("/{OId}/pastes", h.GetOwnerPastes).Methods("GET")

	c.R.HandleFunc("/favicon.ico", faviconHandler)
	c.R.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	apiRouter := c.R.PathPrefix("/api").Subrouter()

	// Pastes
	apiRouter.HandleFunc("/paste", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute}), h.CreatePaste).ServeHTTP).Methods("POST")
	apiRouter.HandleFunc("/paste/{pId}", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute}), h.DeletePaste).ServeHTTP).Methods("DELETE")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}
