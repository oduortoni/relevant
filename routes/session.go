package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"r/db"
)

type Join struct {
	Identifier string
	Session    db.Session
}

func JoinSession(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	sIdentifier := parts[len(parts)-1]
	fmt.Fprintf(w, "<h1>About to join session: %s</h1>", sIdentifier)

	// populating session data from the database
	session, ok := db.SessRetrieve(sIdentifier)
	if !ok {
		fmt.Fprintf(w, "Session not found")
		return
	}
	fmt.Println(session)

	// render session details template
	t := template.Must(template.New("session.html").Funcs(template.FuncMap{
		"getUserName": getUserName,
	}).ParseFiles("./templates/session.html"))

	t.Execute(w, session)
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

func getUserName(userID int) string {
	return "User" + strconv.Itoa(userID)
}
