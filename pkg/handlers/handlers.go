package handlers

import (
	"net/http"

	"github.com/arpushkarev/bookings/pkg/config"
	"github.com/arpushkarev/bookings/pkg/models"
	"github.com/arpushkarev/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	Cfg *config.AppConfig
}

func NewRepo(c *config.AppConfig) *Repository {
	return &Repository{
		Cfg: c,
	}
}

func HandlersRepo(r *Repository) {
	Repo = r
}

func (rep *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["HomeTest"] = "Well Done!"

	render.RenderTemplates(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (rep *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.html", &models.TemplateData{})
}
