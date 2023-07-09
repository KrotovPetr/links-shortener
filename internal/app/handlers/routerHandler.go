package handlers

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/url"
	"net/http"
)

func RouterHandler(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	switch method {
	case http.MethodGet:
		url.RedirectURLHandler(response, request)
	case http.MethodPost:
		url.ShortenURLHandler(response, request)
	default:
		url.ErrorHandler(response, request)
	}

}
