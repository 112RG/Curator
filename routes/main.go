package routes

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/112RG/Curator/repositories"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	files  []string
)

// Run will start the server
func Run() {
	db := connection.setupDatabase()
	userRepo := repositories.NewUserRepo(db)
	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)

	getRoutes()
	router.Run(":5000")
}

func getRoutes() {
	router.Use()
	api := router.Group("/api")
	addPasteRoutes(api)

	web := router.Group("/")
	addWebRoutes(web)
}
