package main

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.RouterHandler)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

}
