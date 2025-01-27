package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Rupesh")
		fmt.Fprint(w, "\nhello")
	})

	fmt.Println("Starting server on port 3000")

	// Enable CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Allow all origins for testing

	// Start the server with CORS handling
	err := http.ListenAndServe(":3000", handlers.CORS(origins, headers, methods)(http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
