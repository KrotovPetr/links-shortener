package main

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/internal/app/url"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {
	router := chi.NewRouter()

	router.Post(`/`, url.ShortenURLHandler)
	router.Get(`/{id}`, url.RedirectURLHandler)

	fmt.Println("Running server on", flagRunAddr)

	return http.ListenAndServe(flagRunAddr, router)
}
