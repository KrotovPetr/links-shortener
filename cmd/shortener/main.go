package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
)

var urlMap = make(map[string]string)

func redirectURLHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id := strings.TrimPrefix(request.URL.Path, "/")
		originalURL, isExist := urlMap[id]

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

func shortenURLHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		body, err := io.ReadAll(request.Body)

		if err != nil {
			http.Error(responseWriter, "Bad Request", http.StatusBadRequest)
			return
		}

		originalURL := string(body)
		shortURL := uuid.New().String()
		urlMap[shortURL] = originalURL

		fullURL := fmt.Sprintf("http://" + request.Host + "/" + shortURL)

		responseWriter.Header().Set("Content-Type", "text/plain")
		responseWriter.WriteHeader(http.StatusCreated)
		responseWriter.Write([]byte(fullURL))
	} else {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Only get method type"))
	}
}

func main() {
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {
	router := chi.NewRouter()

	router.Post(`/`, shortenURLHandler)
	router.Get(`/{id}`, redirectURLHandler)

	fmt.Println("Running server on", flagRunAddr)

	return http.ListenAndServe(flagRunAddr, router)
}
