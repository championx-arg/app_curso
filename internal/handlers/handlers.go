package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/championx-arg/app_curso/internal/config"
	"github.com/championx-arg/app_curso/internal/models"
	"github.com/championx-arg/app_curso/internal/render"
)

//Repo el repositorio usado por handlers
var Repo *Repository

//Repository crea un repositorio nuevo
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHnadlers setea el repositorio para handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home responde ante la solicitud de home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

//Home responde ante la solicitud de home
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

//majors responde ante la solicitud de majors
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

//Generals responde ante la solicitud de Generals
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

//Reservation responde ante la solicitud de Reservation
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

//Availability responde ante la solicitud de Availability
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

//Availability responde ante la solicitud de Availability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start: %s and end %s", start, end)))
	//w.Write([]byte("laconcha tuya"))
}

//cuando defino una estructura que va a ser un json,  se ponen los nombres de las keys con mayuscula ya que deben ser exportables
type jsonResponse struct{ 
	OK bool `json:"ok"`
	Message string `json:"message"`
}


//Availability responde ante la solicitud de Availability
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp:= jsonResponse{
		OK: true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp,"","     ")
	if  err  !=  nil {
		log.Println(err)
	}
	//log.Println(string(out))

	w.Header().Set("Content-Type","application/json")
	w.Write(out)

}

//About responde ante la solicitud de About
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "prueba de envio a templrate"

	//send data to template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
