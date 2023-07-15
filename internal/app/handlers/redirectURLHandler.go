package handlers

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/storage"
	"net/http"
	"strings"
)

func RedirectURLHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id := strings.TrimPrefix(request.URL.Path, "/")
		originalURL, isExist := storage.UrlMap[id]

		if isExist {
			responseWriter.Header().Set("Content-Type", "text/plain")
			responseWriter.Header().Set("Location", originalURL)
			responseWriter.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
		}

	} else {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Only get method type"))
	}

}
