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

// HomeHandler обрабатывает корневой эндпоинт и возвращает HTML из файла index.html
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // Отправляем HTML файл
    http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 1. Парсить html-форму
    err := r.ParseMultipartForm(10 << 20) // 10MB максимум
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusInternalServerError)
        return
    }

    // 2. Получить файл из формы
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // 3. Прочитать данные из файла
    content, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Failed to read file", http.StatusInternalServerError)
        return
    }

    // 4. Передать эти данные в функцию автоопределения из пакета service
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }

    // 5. Создать локальный файл
    // Используем time.Now().UTC().String() для генерации имени файла
    timestamp := time.Now().UTC().String()
    
    // Чтобы получить расширение файла, используйте filepath.Ext()
    ext := filepath.Ext(header.Filename)
    if ext == "" {
        ext = ".txt"
    }
    
    // Очищаем timestamp от недопустимых символов для имени файла
    safeTimestamp := strings.ReplaceAll(strings.ReplaceAll(
        strings.ReplaceAll(timestamp, " ", "_"), 
        ":", "-"), 
        "+", "_")
    
    outputFilename := "converted_" + safeTimestamp + ext

    // 6. Записать в локальный файл результат конвертации строки
    outputFile, err := os.Create(outputFilename)
    if err != nil {
        http.Error(w, "Failed to create output file", http.StatusInternalServerError)
        return
    }
    defer outputFile.Close()

    _, err = outputFile.WriteString(converted)
    if err != nil {
        http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
        return
    }

    // 7. Вернуть результат конвертации строки
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(converted))
}