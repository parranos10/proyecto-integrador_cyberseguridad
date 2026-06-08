package routes

import (
	"net/http"

	"SistemaGestionImagenes/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/images", handlers.GetImages)
	http.HandleFunc("/images/view", handlers.GetImage)
}
