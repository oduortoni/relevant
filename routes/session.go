package routes

import (
	"fmt"
	"net/http"
	"html/template"

	"r/db"
)

func JoinSession(w http.ResponseWriter, r *http.Request) {
}

func ListSessions(w http.ResponseWriter, r *http.Request) {
	sessions, ok := db.SessList()
	if !ok {
		fmt.Fprintf(w, "No sessions to display")
		return
	}
	fmt.Println(sessions)

	t := template.Must(template.ParseFiles("./templates/sessions.html"));
	t.Execute(w, sessions)
}
