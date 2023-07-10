package main

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/url"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Post(`/`, url.ShortenURLHandler)
	router.Get(`/{id}`, url.RedirectURLHandler)

	err := http.ListenAndServe(`:8080`, router)
	if err != nil {
		panic(err)
	}

}
