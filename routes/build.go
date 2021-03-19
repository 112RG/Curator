package routes

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/112RG/Curator/db"
	"github.com/112RG/Curator/repositories"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	router          = gin.New()
	files           []string
	pasteRepository *repositories.PasteRepo
)

// Run will start the server
func Build() *gin.Engine {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	db := db.ConnectDB()
	pasteRepository = repositories.NewPasteRepo(db)

	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)
	router.Use(logger.SetLogger())
	//router.Use(favicon.New("./favicon.png"))
	router.Use(static.Serve("/assets", static.LocalFile("./assets", false)))
	getRoutes()
	return router
}

func getRoutes() {
	router.Use()
	api := router.Group("/api")
	addPasteRoutes(api)

	web := router.Group("/")
	addWebRoutes(web)
}
