package routes

import (
	"fmt"
	"net/http"

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
	fmt.Fprintf(w, "%v", sessions)
}
