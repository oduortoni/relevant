package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"r/db"
)

func GLogin(w http.ResponseWriter, r *http.Request, path string) {
	// fetch the templates (GET request)
	if r.URL.Path != "/login" {
		ServeError(w, "Page Not Found", http.StatusNotFound, path)
		return
	}

	// execute the template
	if r.Method == "GET" {
		t, err := template.ParseFiles(path)
		if err != nil {
			ServeError(w, "Internal Server Error", http.StatusInternalServerError, path)
			return
		}
		t.ExecuteTemplate(w, "templates/login.html", nil)
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
