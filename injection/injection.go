package injection

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/112RG/Curator/db"
	"github.com/112RG/Curator/handler"
	"github.com/112RG/Curator/repositories"
	"github.com/112RG/Curator/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thinkerou/favicon"
)

var (
	files []string
)

func Inject() (*gin.Engine, error) {
	db := db.ConnectDB()

	pasteRepository := repositories.NewPasteRepository(db)
	pasteService := service.NewPasteService(&service.USConfig{
		PasteRepository: pasteRepository,
	})
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

	router := gin.New()

	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)
	router.Use(logger.SetLogger())
	router.Use(favicon.New("./favicon.png"))
	router.Use(static.Serve("/assets", static.LocalFile("./assets", false)))
	handler.NewHandler(&handler.Config{
		R:            router,
		PasteService: pasteService,
	})

	return router, nil
}
