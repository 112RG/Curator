package routes

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/112RG/Curator/db"
	"github.com/112RG/Curator/repositories"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

var (
	router          = gin.Default()
	files           []string
	pasteRepository *repositories.PasteRepo
)

// Run will start the server
func Run() {
	db := db.ConnectDB()
	pasteRepository = repositories.NewPasteRepo(db)

	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)
	router.Use(static.Serve("/assets", static.LocalFile("./assets", false)))
	router.Use(favicon.New("./favicon.ico"))
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
