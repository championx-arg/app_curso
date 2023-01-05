package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/championx-arg/app_curso/pkg/config"
	"github.com/championx-arg/app_curso/pkg/handlers"
	"github.com/championx-arg/app_curso/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"
var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {

	//cambiar esto a true cuando se este en produccion

	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("no se puede crear template cache")
	}

	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/",handlers.Repo.Home)
	//http.HandleFunc("/about",handlers.Repo.About)

	//http.ListenAndServe(portNumber,nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	fmt.Printf("Starting server on %s", portNumber)

}
