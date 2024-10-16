package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"r/db"
)

func JoinSession(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	sIdentifier := parts[len(parts)-1]
	fmt.Fprintf(w, "<h1>About to join session: %s</h1>", sIdentifier)
}

func ListSessions(w http.ResponseWriter, r *http.Request) {
	sessions, ok := db.SessList()
	if !ok {
		fmt.Fprintf(w, "No sessions to display")
		return
	}
	fmt.Println(sessions)

	t := template.Must(template.ParseFiles("./templates/sessions.html"))
	t.Execute(w, sessions)
}
