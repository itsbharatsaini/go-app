package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, Inc42")
}

func main() {
	// Register a handler for the root URL path "/"
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
