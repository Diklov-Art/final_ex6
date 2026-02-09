package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
    // Создаем логгер
    logger := log.New(os.Stdout, "MORSE-CONVERTER: ", log.Ldate|log.Ltime|log.Lshortfile)
    
    
    srv := server.New(logger)
    
    
    go func() {
        if err := srv.Start(); err != nil && err != http.ErrServerClosed {
            
            logger.Fatalf("Server failed to start: %v", err)
        }
    }()
    
    logger.Println("Server started on http://localhost:8080")
    logger.Println("Press Ctrl+C to stop the server")
    
    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    logger.Println("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        logger.Fatalf("Server forced to shutdown: %v", err)
    }
    
    logger.Println("Server stopped")
}