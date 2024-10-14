package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"r/db"
)

func GLogin(w http.ResponseWriter, r *http.Request) {
	// execute the template
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/login.html")
		if err != nil {
			ServeError(w, "Internal Server Error", http.StatusInternalServerError, "./templates/error.html")
			return
		}
		t.Execute(w, "templates/login.html")
	}
}

func PLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, success := db.CRetrieve(username)
	if !success {
		fmt.Fprintf(w, "User not found")
		return
	}
	if user.Password == password {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Fprintf(w, "Invalid credentials")
	}
}
