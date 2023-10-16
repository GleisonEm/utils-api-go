package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gleisonem/utils-api-go/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	r := mux.NewRouter()
	r.HandleFunc("/tiktok/video/base64", controllers.getBase64ByVideo).Methods("GET")

	http.Handle("/", r)

	address := ":" + port
	fmt.Printf("Servidor escutando na porta %s...\n", port)
	http.ListenAndServe(address, nil)
}
