package main

import (
	"fmt"
	"net/http"
	"os"

	"r/db"
	"r/routes"
)

func main() {
	if len(os.Args) == 2 {
		initializeSessionDB()
	}
	// serve static files
	http.HandleFunc("/static/", routes.Static)

	// ordinary routes
	http.HandleFunc("/", routes.Homepage)
	http.HandleFunc("/gregister", routes.GRegister)
	http.HandleFunc("/pregister", routes.PRegister)
	http.HandleFunc("/glogin", routes.GLogin)
	http.HandleFunc("/plogin", routes.PLogin)

	http.HandleFunc("/gsessions", routes.ListSessions)
	http.HandleFunc("/sjoin/", routes.JoinSession)

	fmt.Println("Server running on :9000")
	http.ListenAndServe(":9000", nil)
}

func initializeSessionDB() {
	u1 := db.User{
		Name:     "uone",
		Password: "uone",
		Id:       1,
	}
	u2 := db.User{
		Name:     "utwo",
		Password: "utwo",
		Id:       2,
	}
	u3 := db.User{
		Name:     "utres",
		Password: "utres",
		Id:       3,
	}
	s1 := db.Session{
		Members:    []db.User{u1, u2},
		OwnerId:    1,
		Identifier: "session uno",
	}

	s2 := db.Session{
		Members:    []db.User{u1, u2, u3},
		OwnerId:    1,
		Identifier: "session dos",
	}

	db.SessSave(s1.Identifier, s1)
	db.SessSave(s2.Identifier, s2)

	s, ok := db.SessRetrieve(s1.Identifier)
	if !ok {
		fmt.Println("No such session")
	} else {
		fmt.Println(s)
	}

	sessions, ok := db.SessRetrieve(s2.Identifier)
	if !ok {
		fmt.Println("No such session")
	} else {
		fmt.Println(sessions)
	}
}
