package handlers

import (
	"encoding/json"
	"github.com/KrotovPetr/links-shortener.git/internal/app/storage"
	"io"
	"log"
	"net/http"
)

type URLData struct {
	URL string `json:"url"`
}

type ResultData struct {
	RESULT string `json:"result"`
}

func ReturnShortenURLHandler(responseWriter http.ResponseWriter, request *http.Request) {
	body, _ := io.ReadAll(request.Body)

	var result URLData

	if err := json.Unmarshal(body, &result); err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Failed to parse request body as JSON"))
		return
	}

	shortenURL := storage.URLMap[result.URL]

	var resultJSON = ResultData{
		RESULT: shortenURL,
	}

	responseJSON, err := json.Marshal(resultJSON)
	if err != nil {
		log.Println("Failed to marshal error data as JSON:", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write([]byte(responseJSON))
}
