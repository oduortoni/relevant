package routes

import (
	"fmt"
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	filepath := "." + r.URL.Path

	// file statistics
	stats, err := os.Stat(filepath)
	if err != nil {
		fmt.Fprintf(w, "<h1>File %s does not exist</h1>", filepath)
		return
	}

	if stats.IsDir() {
		fmt.Fprintf(w, "<h1>Permisiion denied</h1>")
		return
	}

	http.ServeFile(w, r, filepath)
}
