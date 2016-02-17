package main

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Hello)
	
	http.ListenAndServe(":8080", mux)
}
