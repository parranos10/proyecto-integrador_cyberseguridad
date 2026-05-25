package routes

import (
	"net/http"

	"SistemaGestionImagenes/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/upload", handlers.UploadImage)

	http.HandleFunc("/images", handlers.GetImages)
}
