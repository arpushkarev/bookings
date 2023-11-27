package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arpushkarev/bookings/internal/config"
	"github.com/arpushkarev/bookings/internal/models"
	"github.com/arpushkarev/bookings/internal/render"
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

	render.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (rep *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (rep *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (rep *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (rep *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (rep *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s, end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (rep *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (rep *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
