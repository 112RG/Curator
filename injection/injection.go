package injection

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"curator/db"
	"curator/handler"
	"curator/repositories"
	"curator/service"

	"github.com/gorilla/mux"
	"github.com/michaeljs1990/sqlitestore"
)

// Command line flags.
var (
	dsn = flag.String("dsn", "", "datasource name")
)

func Inject() (*mux.Router, error) {
	flag.Parse()
	db := db.ConnectDB(dsn)

	pasteRepository := repositories.NewPasteRepository(db)
	pasteService := service.NewPasteService(&service.USConfig{
		PasteRepository: pasteRepository,
	})

	sessionStore, err := sqlitestore.NewSqliteStore("./sessions.db", "sessions", "/", 3600, []byte("<SecretKey>"))
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()

	templates := ParseTemplates()
	handler.NewHandler(&handler.Config{
		R:               router,
		PasteService:    pasteService,
		TemplateService: templates,
		SessionService:  sessionStore,
	})

	return router, nil
}

func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
