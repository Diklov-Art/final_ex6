package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура сервера
type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

// New создает новый сервер
func New(logger *log.Logger) *Server {
	// Создаем роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем HTTP сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

// Start запускает сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully останавливает сервер
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
