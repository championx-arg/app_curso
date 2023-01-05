package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/championx-arg/app_curso/pkg/config"
	"github.com/championx-arg/app_curso/pkg/models"
)

/*
var functions = template.FuncMap{

}
*/

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	// esto es para decirle al programa si voy a usar o no
	//el generador de cache, esto sirve para que actualice en cada
	//llamada
	if app.UseCache {
		//tc, err := CreateTemplateCache()
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// create a template cache
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("no se pudo crear el template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td) // esto es para agregar informacion por defecto a algunas paginas

	_ = t.Execute(buf, td)
	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")

	fmt.Printf("path error %s\n", err)


	files := path.Dir("././")
    fmt.Println(files)

	

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// ESTA ES LA VERSION INICIAL. MAS SIMPLE PERO HAY QUE AGREGAR CADA TEMPLATE O ARCHIVO NUEVO, LA VERSION NUEVA SE HACE SOLO
/*
package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)


func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")

	err:= parsedTemplate.Execute(w,nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error

	//quiero ver si el template ya esta en memria o no
	_, inMap := templateCache[t] //devuelve el contenido de templatecache y un true o false si tiene el contenido.
	if !inMap {
		//tengo que crear la template ya que no esta
		log.Println("creanto template en cache")
		err = CreateTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	}else {
		//esta en memoria
		log.Println("usando template desde cache")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache(t string) error {

	//aca irian las layauts que se van a incluir
	templates := []string{
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.tmpl",
	}

	//parsear el template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache
	templateCache[t] = tmpl

	return nil

}

*/
