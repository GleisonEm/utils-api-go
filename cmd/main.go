package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gleisonem/utils-api-go/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	r.HandleFunc("/youtube/video/base64", controllers.GetBase64ByVideo).Methods("GET")
	r.HandleFunc("/youtube/video/base642", controllers.GetBase64ByVideo2).Methods("GET")

	http.Handle("/", r)

	address := ":" + port
	fmt.Printf("Servidor escutando na porta %s...\n", port)
	http.ListenAndServe(address, nil)
}
