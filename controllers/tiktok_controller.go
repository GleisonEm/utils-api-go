// controllers/search_controller.go

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"encoding/base64"
	"io"
)

type Video struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

func getBase64ByVideo(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("url")
	if name == "" {
		http.Error(w, "Parâmetro 'name' não especificado na URL", http.StatusBadRequest)
		return
	}

	videoData, err := getVideoNoWM(url)
	if err != nil {
		http.Error(w, "Erro ao obter a url do vídeo", http.StatusInternalServerError)
		return
	}

	base64String, err := getBase64Video(videoData.URL)
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

func getVideoNoWM(url string) (*Video, error) {
	id, err := getIdVideo(url)
	if err != nil {
		return nil, err
	}
	apiURL := fmt.Sprintf("https://api16-normal-c-useast1a.tiktokv.com/aweme/v1/feed/?aweme_id=%s", id)
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", "TikTok 26.2.0 rv:262018 (iPhone; iOS 14.4.2; en_US) Cronet")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	list := data["aweme_list"].([]interface{})
	video := list[0].(map[string]interface{})
	playAddr := video["video"].(map[string]interface{})["play_addr"].(map[string]interface{})
	urlList := playAddr["url_list"].([]interface{})
	urlMedia := urlList[0].(string)

	return &Video{URL: urlMedia, ID: id}, nil
}

func getIdVideo(url string) (string, error) {
	if !strings.Contains(url, "/video/") {
		return "", fmt.Errorf("invalid URL")
	}
	id := strings.Split(url, "/video/")[1]
	if len(id) > 19 {
		id = strings.Split(id, "?")[0]
	}
	return id, nil
}

func getBase64Video(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var base64String string

	encoder := base64.NewEncoder(base64.StdEncoding, &base64String)
	defer encoder.Close()

	_, err = io.Copy(encoder, response.Body)
	if err != nil {
		return "", err
	}

	return base64String, nil
}
