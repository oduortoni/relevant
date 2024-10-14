package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"r/db"
)

func PRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, exists := db.CRetrieve(username)
	if exists {
		fmt.Fprintf(w, "A user with the name, %s, already exists\n", username)
		return
	}
	
	user := db.User{
		Name: username,
		Password: password,
	}

	db.CSave(username, user)

	fmt.Fprintf(w, "Successfully saved\n\n\tuser: %s\n\tPass: %s\n", username, password)
}

func GRegister(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/register.html"))
	t.Execute(w, nil)
}
