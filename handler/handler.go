package handler

import (
	"github.com/112RG/Curator/model"
	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	PasteService model.PasteService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R            *gin.Engine
	PasteService model.PasteService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		PasteService: c.PasteService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group("/")
	g.GET("/:id", h.Get)

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
