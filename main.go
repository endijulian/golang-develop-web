package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/endi", endiHandler)

	log.Println("Startting web on Port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, Saya sedang belajar golang web"))
}

func endiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Nama saya Endi Julian"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to page HOME"))
}
