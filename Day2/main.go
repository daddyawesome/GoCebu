package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "🦄 Hello World! - Go Web App")
	})

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}