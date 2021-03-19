package main

import (
	"github.com/112RG/Curator/db"
	"github.com/112RG/Curator/repositories"
	"github.com/112RG/Curator/service"
	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/memrizr/account/handler"
)

func Inject() (*gin.Engine, error) {
	db := db.ConnectDB()

	pasteRepository := repositories.NewPasteRepo(db)

	pasteService := service.NewPasteService(&service.USConfig{
		PasteRepository: pasteRepository,
	})
	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R:            router,
		PasteService: pasteService,
	})
	return router, nil
}
