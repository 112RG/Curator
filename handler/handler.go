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
	} // currently has no properties
	// Create an account group
	//url := ginSwagger.URL("http://localhost:5000/swagger/doc.json") // The url pointing to API definition
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

	/* 	i := c.R.Group("/")

	   	i.GET("/:pId", h.GetPaste)
	   	i.GET("/:pId/raw", h.GetPasteRaw)
	   	i.GET("/:pId/owners", h.TestPaste)

	   	i.GET("/", h.GetIndex)

	   	sw := c.R.Group("/swagger")
	   	sw.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	   	v1 := c.R.Group("/v1")
	   	p := v1.Group("/paste")
	   	p.POST("/", h.CreatePaste)
	   	p.DELETE("/:pId", h.DeletePaste)
	*/
	/* 	if gin.Mode() != gin.TestMode {
	   		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
	   		g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
	   		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
	   		g.PUT("/details", middleware.AuthUser(h.TokenService), h.Details)
	   		g.POST("/image", middleware.AuthUser(h.TokenService), h.Image)
	   		g.DELETE("/image", middleware.AuthUser(h.TokenService), h.DeleteImage)
	   	} else {
	   		g.GET("/", h.Me)

	   	} */

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "relative/path/to/favicon.ico")
}
