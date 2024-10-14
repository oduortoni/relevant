package main

import (
	"net/http"

	"r/routes"
)

func main() {
	http.ListenAndServe(":9000", nil)

	http.HandleFunc("/", routes.Homepage)
}
