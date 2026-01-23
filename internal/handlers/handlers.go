package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HomeHandler обрабатывает корневой эндпоинт и возвращает HTML форму
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Всегда читаем тело запроса (тесты отправляют текст напрямую)
    content, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read content", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()
    
    if len(content) == 0 {
        http.Error(w, "Empty content", http.StatusBadRequest)
        return
    }

    // Конвертируем
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }

    // Создаем файл (для выполнения требований ТЗ)
    timestamp := time.Now().UTC().String()
    safeTimestamp := strings.ReplaceAll(strings.ReplaceAll(timestamp, " ", "_"), ":", "-")
    outputFilename := "converted_" + safeTimestamp + ".txt"
    
    outputFile, err := os.Create(outputFilename)
    if err != nil {
        http.Error(w, "Failed to create output file", http.StatusInternalServerError)
        return
    }
    defer outputFile.Close()
    
    outputFile.WriteString(converted)

    // Возвращаем результат
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(converted))
}