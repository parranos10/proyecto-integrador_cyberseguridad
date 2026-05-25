package handlers

import (
	"fmt"
	"net/http"

	"SistemaGestionImagenes/services"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "No se pudo obtener la imagen", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = services.SaveImage(file, handler.Filename)
	if err != nil {
		http.Error(w, "Error al guardar la imagen", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Imagen subida y guardada correctamente")
}
