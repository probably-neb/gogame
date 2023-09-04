package main

import (
	"log"
	"net/http"
)

func main() {
	// Set routing rules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
