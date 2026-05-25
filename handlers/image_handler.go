package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func GetImages(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Aquí se mostrarán las imágenes guardadas"))
}
