package main

import (
	"fmt"
	"net/http"
	"os"

	"r/db"
	"r/routes"
)

func main() {
	_, err := os.Stat("./database")
	if os.IsNotExist(err) {
		err := os.Mkdir("database", 0755) // Use desired permissions, e.g., 0755
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Database directory created.")
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		return
	} else {;}

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
		Name:     "mikaela",
		Password: "mika",
		Id:       1,
	}
	u2 := db.User{
		Name:     "rafael",
		Password: "rafa",
		Id:       2,
	}
	u3 := db.User{
		Name:     "rodrigo",
		Password: "rodri",
		Id:       3,
	}
	u4 := db.User{
		Name:     "angel",
		Password: "angel",
		Id:       4,
	}

	msg101 := db.Message{
		MemberId: 1,
		Content: "Hello leute!",
	}
	msg102 := db.Message{
		MemberId: 2,
		Content: "kool",
	}
	msg103 := db.Message{
		MemberId: 3,
		Content: "Mhh!",
	}
	msg104 := db.Message{
		MemberId: 1,
		Content: "Sq, welcome to the chat. Today we will be talking about",
	}
	s1 := db.Session{
		Members:    []db.User{u1, u2, u2},
		Messages: []db.Message{msg101, msg102, msg103, msg104},
		OwnerId:    1,
		Identifier: "session uno",
	}

	msg201 := db.Message{
		MemberId: 1,
		Content: "Can someone tell me what a diode is?",
	}
	msg202 := db.Message{
		MemberId: 2,
		Content: "What! I had the exact same question in mind",
	}
	msg203 := db.Message{
		MemberId: 3,
		Content: "I could try to...",
	}
	msg204 := db.Message{
		MemberId: 1,
		Content: "Please do so rodri!",
	}
	msg205 := db.Message{
		MemberId: 3,
		Content: "A diode is sort of a one way valve",
	}
	msg206 := db.Message{
		MemberId: 4,
		Content: "Tah makes no sense, rodri. Dont assume everyone knows what a valve is.",
	}
	s2 := db.Session{
		Members:    []db.User{u1, u2, u3, u4},
		Messages: []db.Message{msg201, msg202, msg203, msg204, msg205, msg206},
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
