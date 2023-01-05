package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//appconfig tiene la configuracion de la aplicacion
//aca irian variables  y datos disponibles para toda la aplicacion
//vendria siendo como las variables globales, pero es una estructura
//donde estan todas las variables que quiero que sean globales, entonces
//solo tengo que importar AppConfig en los archivos, a esto se lo llama
//repositorio...

type AppConfig struct{
	UseCache bool
	TemplateCache map[string]*template.Template
	InProduction bool
	Session *scs.SessionManager

}