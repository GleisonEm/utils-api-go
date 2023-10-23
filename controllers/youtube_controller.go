package controllers

import (
	"encoding/base64"
	"io"
	"net/http"
	"encoding/json"
	"bytes"
	"os"
)

func GetBase64ByVideo(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Parâmetro 'url' não especificado na URL", http.StatusBadRequest)
		return
	}

	base64String, err := getBase64Video(url)
	if err != nil {
		http.Error(w, "Erro ao obter o vídeo em base64", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"videoBase64": base64String,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func GetBase64ByVideo2(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Parâmetro 'url' não especificado na URL", http.StatusBadRequest)
		return
	}

	getBase64Video2(url)

	response := map[string]string{
		"videoBase64": "deu certo",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}


	func getBase64Video2(url string) {
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		file, err := os.Create("tiktokvideo_base64.txt") // Nome do arquivo de saída
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := base64.NewEncoder(base64.StdEncoding, file)
		defer encoder.Close()

		_, err = io.Copy(encoder, response.Body)
		if err != nil {
			panic(err)
		}

		println("Vídeo convertido para base64 e salvo com sucesso.")
	}

	func getBase64Video(url string) (string, error) {
		response, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer response.Body.Close()

		buf := &bytes.Buffer{}
		encoder := base64.NewEncoder(base64.StdEncoding, buf)
		defer encoder.Close()

		_, err = io.Copy(encoder, response.Body)
		if err != nil {
			return "", err
		}

		base64String := buf.String()
		return base64String, nil
	}