package main

import (
	"fmt"
	"net/http"

	"r/routes"
)

func main() {
	http.HandleFunc("/", routes.Homepage)
	http.HandleFunc("/gregister", routes.GRegister)
	http.HandleFunc("/pregister", routes.PRegister)
	// http.HandleFunc("/glogin", routes.GLogin)
	// http.HandleFunc("/plogin", routes.PLogin)

	fmt.Println("Server running on :9000")
	http.ListenAndServe(":9000", nil)
}
