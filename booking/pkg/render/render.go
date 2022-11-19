package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/artmadar/golang-monolith-web-app/booking/pkg/config"
	"github.com/artmadar/golang-monolith-web-app/booking/pkg/models"
)

var functions = template.FuncMap{}
var templatesFolderPath = "../../templates"

// var templateCache map[string]*template.Template

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// if templateCache == nil {
	// 	cache, err := GetTemplateCache()
	// 	if err != nil {
	// 		log.Println("Error getting template cache: ", err)
	// 	}

	// 	templateCache = cache
	// }

	var tc map[string]*template.Template
	if app.UserCache {
		tc = app.TempateCache
	} else {
		var err error
		tc, err = GetTemplateCache()

		if err != nil {
			log.Fatal("Template not exist or it was imposible to get it from cache")
		}
	}

	parsedTemplate, ok := tc[tmpl]

	if !ok {
		log.Fatal("Template not exist or it was imposible to get it from cache")
	}

	td = AddDefaultTemplateData(td)

	err := parsedTemplate.Execute(w, td)

	if err != nil {
		log.Println("Error parsing template: ", err)
	}
}

func GetTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(templatesFolderPath + "/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Error parsing template: ", err)
			return myCache, err
		}

		match, err := filepath.Glob(templatesFolderPath + "/*.layout.html")
		if err != nil {
			log.Println("Error Glob layout: ", err)
			return myCache, err
		}

		if len(match) > 0 {
			ts, err = ts.ParseGlob(templatesFolderPath + "/*.layout.html")
			if err != nil {
				log.Println("Error ParseGlob layout: ", err)
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
