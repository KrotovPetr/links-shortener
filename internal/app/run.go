package app

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Run(flagRunAddr string) error {
	router := chi.NewRouter()

	router.Post(`/`, handlers.ShortenURLHandler)
	router.Get(`/{id}`, handlers.RedirectURLHandler)

	return http.ListenAndServe(flagRunAddr, router)
}
