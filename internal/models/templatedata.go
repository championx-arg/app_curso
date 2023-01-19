package models

//esta estructura es para poder enviar datos a los templates. Se crean varios tipos de 
//variables para tener opciones de que se va a enviar.
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string //success message por ejemplo
	Warning string //warning message por ejemplo
	Error string //error message por ejemplo
}