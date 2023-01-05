package handlers

import (
	"net/http"
	"github.com/championx-arg/app_curso/pkg/config"
	"github.com/championx-arg/app_curso/pkg/models"
	"github.com/championx-arg/app_curso/pkg/render"
)

//Repo el repositorio usado por handlers
var Repo *Repository

//Repository crea un repositorio nuevo
type Repository struct {

	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHnadlers setea el repositorio para handlers
func NewHandlers(r *Repository){
	Repo = r
}

//Home responde ante la solicitud de home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{} )
}

//About responde ante la solicitud de home
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "prueba de envio a templrate"

	//send data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}