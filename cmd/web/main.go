package main

import (
	"fmt"
	"net/http"
	"udemygo/pkg/handlers"
)

const port = ":8080"

// main is the main function that starts the HTTP server.
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Listening at", port)
	http.ListenAndServe(port, nil)
}
