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

    var content []byte
    var err error

    // Сначала пытаемся прочитать как multipart форму
    err = r.ParseMultipartForm(10 << 20)
    if err == nil {
        // Это multipart форма
        file, header, err := r.FormFile("file")
        if err != nil {
            // Если нет файла, читаем тело
            content, err = io.ReadAll(r.Body)
        } else {
            defer file.Close()
            content, err = io.ReadAll(file)
            
            // Сохраняем оригинальное имя файла для расширения
            if header != nil {
                // Используем расширение оригинального файла
                ext := filepath.Ext(header.Filename)
                if ext == "" {
                    ext = ".txt"
                }
                
                // Создаем файл с результатом
                timestamp := time.Now().UTC().String()
                safeTimestamp := strings.ReplaceAll(strings.ReplaceAll(timestamp, " ", "_"), ":", "-")
                outputFilename := "converted_" + safeTimestamp + ext
                
                outputFile, err := os.Create(outputFilename)
                if err == nil {
                    defer outputFile.Close()
                    // Конвертируем и записываем
                    converted, _ := service.Convert(string(content))
                    outputFile.WriteString(converted)
                    
                    // Возвращаем результат
                    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
                    w.WriteHeader(http.StatusOK)
                    w.Write([]byte(converted))
                    return
                }
            }
        }
    } else {
        // Если не multipart, читаем тело запроса
        content, err = io.ReadAll(r.Body)
    }
    
    if err != nil {
        http.Error(w, "Failed to read content", http.StatusInternalServerError)
        return
    }
    
    // Конвертируем
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }
    
    // Создаем файл с результатом
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