package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a file server handler for the "public" directory.
	fs := http.FileServer(http.Dir("./public"))

	// Register the handler to serve files from "public" when requests come to "/assets/".
	// http.StripPrefix removes "/assets/" from the URL path before passing it to fs.
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("Serving static content on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
