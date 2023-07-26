package app

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/handlers"
	logger2 "github.com/KrotovPetr/links-shortener.git/internal/app/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

func Run(flagRunAddr string) error {
	logger, _ := zap.NewProduction(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(logger2.LoggingMiddleware(logger))

	router.Post(`/`, handlers.ShortenURLHandler)
	router.Post(`/api/shorten`, handlers.ReturnShortenURLHandler)

	router.Get(`/{id}`, handlers.RedirectURLHandler)

	return http.ListenAndServe(flagRunAddr, router)
}
