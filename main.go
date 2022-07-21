package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/endi", endiHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	})

	log.Println("Startting web on Port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to page HOME"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, Saya sedang belajar golang web"))
}

func endiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Nama saya Endi Julian"))
}
