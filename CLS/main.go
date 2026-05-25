package main

import (
	"fmt"
	"net/http"

	"SistemaGestionImagenes/routes"
)

func main() {

	routes.SetupRoutes()

	fmt.Println("Servidor iniciado en puerto 8080")

	http.ListenAndServe(":8080", nil)
}
