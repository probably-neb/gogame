package main

import (
	"log"
	"net/http"
    "fmt"
)


func main() {
    var count int = 0;
	// Set routing rules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
    http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
        count++
        fmt.Fprintf(w, "%d", count)
    })

    fmt.Println("Server started at 8080 port")
	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
