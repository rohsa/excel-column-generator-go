package rendoring

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("view.html"))

// RenderTemplate renders the given template on the web UI
// It pre-assumes that the given template is of type html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
