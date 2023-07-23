package handlers

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/internal/app/storage"
	"github.com/KrotovPetr/links-shortener.git/internal/app/utils"
	"io"
	"net/http"
)

func ShortenURLHandler(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(responseWriter, "Bad Request", http.StatusBadRequest)
		return
	}

	originalURL := string(body)
	shortURL := utils.ShortLinksCreator()

	storage.URLMap[shortURL] = originalURL

	fullURL := fmt.Sprintf("http://" + request.Host + "/" + shortURL)

	responseWriter.Header().Set("Content-Type", "text/plain")
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write([]byte(fullURL))
}
