package handlers

import (
	"net/http"

	"github.com/artmadar/golang-monolith-web-app/pkg/config"
	"github.com/artmadar/golang-monolith-web-app/pkg/models"
	"github.com/artmadar/golang-monolith-web-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplates(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {

	stringMap := make(map[string]string)

	stringMap["test"] = "hello again "

	render.RenderTemplates(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
