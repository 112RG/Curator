package handler

import (
	"net/http"
	"text/template"

	"github.com/112RG/Curator/model"
	"github.com/gorilla/mux"
)

// Handler struct holds required services for handler to function
type Handler struct {
	PasteService    model.PasteService
	TemplateService *template.Template
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *mux.Router
	PasteService    model.PasteService
	TemplateService *template.Template
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {

	// Create a handler (which will later have injected services)
	h := &Handler{
		PasteService:    c.PasteService,
		TemplateService: c.TemplateService,
	}

	c.R.HandleFunc("/", h.GetIndex).Methods("GET")

	c.R.HandleFunc("/{pId}", h.GetPaste).Methods("GET")
	c.R.HandleFunc("/{pId}/raw", h.GetPasteRaw).Methods("GET")
	c.R.HandleFunc("/{OId}/pastes", h.TestPaste).Methods("GET")

	c.R.HandleFunc("/favicon.ico", faviconHandler)
	c.R.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	c.R.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/"))))

	apiRouter := c.R.PathPrefix("/api").Subrouter()
	// Pastes
	apiRouter.HandleFunc("/paste", h.CreatePaste).Methods("POST")
	apiRouter.HandleFunc("/paste/{pId}", h.DeletePaste).Methods("DELETE")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
