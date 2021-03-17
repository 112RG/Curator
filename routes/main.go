package routes

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	files  []string
)

// Run will start the server
func Run() {
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
	api := router.Group("/api")
	addPasteRoutes(api)

	web := router.Group("/")
	addWebRoutes(web)
}
