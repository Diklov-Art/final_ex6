package server

import (
    "context"
    "log"
    "net/http"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура сервера с полями для логгера и http-сервера
type Server struct {
    logger     *log.Logger
    httpServer *http.Server
}

// New создает новый сервер
func New(logger *log.Logger) *Server {
    // Создаем http-роутер
    router := http.NewServeMux()

    // Регистрируем хендлеры в http-роутере
    router.HandleFunc("/", handlers.HomeHandler)
    router.HandleFunc("/upload", handlers.UploadHandler)

    // Создаем экземпляр структуры http.Server
    httpServer := &http.Server{
        Addr:         ":8080",      // используем порт 8080
        Handler:      router,       
        ErrorLog:     logger,       
        ReadTimeout:  5 * time.Second,   // таймаут для чтения. 5 секунд
        WriteTimeout: 10 * time.Second,  // таймаут для записи. 10 секунд
        IdleTimeout:  15 * time.Second,  // таймаут ожидания следующего запроса. 15 секунд
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