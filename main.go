package main

import (
	// "fmt"
	// "net/http"

	// "r/routes"
	"fmt"
	"log"

	"r/db"
)

func main() {
	// http.HandleFunc("/", routes.Homepage)

	// fmt.Println("Server running on :9000")
	// http.ListenAndServe(":9000", nil)

	// db.CSave("joe", db.User{
	// 	Name:     "Joe",
	// 	Password: "joe",
	// })

	user, found := db.CRetrieve("joe")
	if !found {
		log.Fatalln("User not found")
	} else {
		fmt.Println(user)
	}
}
