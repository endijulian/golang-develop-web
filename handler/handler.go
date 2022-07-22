package handler

import (
	"GolangDevelopWeb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to page HOME"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Endi Julian",
	// 	"content": "Saya sedang belajar golang",
	// 	"Umur":    20,
	// }

	// data := entity.Product{Id: 1, Nama: "Honda Civic", Price: 4000000000, Stock: 3}
	data := []entity.Product{
		{Id: 1, Nama: "Honda Civic", Price: 4000000000, Stock: 11},
		{Id: 2, Nama: "Honda Jazz", Price: 4000000000, Stock: 8},
		{Id: 3, Nama: "Lamborghini", Price: 4000000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, Saya sedang belajar golang web"))
}

func EndiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Nama saya Endi Julian"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Page Product"))
	// fmt.Fprintf(w, "Product Page : %d", idNumb)

	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
		return
	}
}
