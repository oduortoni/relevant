package routes

import (
	"net/http"
	"html/template"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/homepage.html"))
	t.Execute(w, nil)
}
