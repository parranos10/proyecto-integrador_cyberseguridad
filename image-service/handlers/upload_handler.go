package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"SistemaGestionImagenes/services"
)

func Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error al obtener la imagen", http.StatusBadRequest)
		return
	}
	defer file.Close()

	extension := filepath.Ext(handler.Filename)
	uniqueName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	savePath := filepath.Join("./uploads", uniqueName)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error al leer los datos de la imagen", http.StatusInternalServerError)
		return
	}

	bytesCifrados, err := services.EncryptAESCBC(fileBytes, services.ClaveSecreta)
	if err != nil {
		http.Error(w, "Error al cifrar la imagen", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(savePath, bytesCifrados, 0644)
	if err != nil {
		http.Error(w, "Error al guardar la imagen cifrada", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Imagen subida con éxito", "filename": "%s"}`, uniqueName)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Servidor Core de Gestión de Imágenes en Go activo"}`)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	files, err := os.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Error al leer la carpeta de imágenes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"images": [`)
	for i, file := range files {
		if !file.IsDir() {
			fmt.Fprintf(w, `"%s"`, file.Name())
			if i < len(files)-1 {
				fmt.Fprintf(w, ", ")
			}
		}
	}
	fmt.Fprintf(w, `]}`)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		http.Error(w, "Falta el parámetro 'name'", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("./uploads", fileName)

	cipherBytes, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusNotFound)
		return
	}

	plainBytes, err := services.DecryptAESCBC(cipherBytes, services.ClaveSecreta)
	if err != nil {
		http.Error(w, "Error al descifrar la imagen", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(plainBytes)
}
