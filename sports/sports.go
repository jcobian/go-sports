package main

import (
	"fmt"
	"net/http"
)

func playersHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/players/"):]
	fmt.Fprintf(w, "<h1>%s</h1>", title)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yoohoo. Big summer blowout: %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/players/", playersHandler)
	http.ListenAndServe(":8080", nil)
}
