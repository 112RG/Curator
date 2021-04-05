package injection

import (
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
)

func Inject() (*mux.Router, error) {
	db := db.ConnectDB()

	pasteRepository := repositories.NewPasteRepository(db)
	pasteService := service.NewPasteService(&service.USConfig{
		PasteRepository: pasteRepository,
	})

	router := mux.NewRouter()

	templates := ParseTemplates()
	handler.NewHandler(&handler.Config{
		R:               router,
		PasteService:    pasteService,
		TemplateService: templates,
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
