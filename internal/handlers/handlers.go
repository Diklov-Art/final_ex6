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
    var filename string
    var err error

    // Проверяем, multipart ли это (форма с файлом из браузера)
    if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
        // Парсим html-форму
        err = r.ParseMultipartForm(10 << 20) // 10MB максимум
        if err != nil {
            http.Error(w, "Failed to parse form", http.StatusInternalServerError)
            return
        }

        // Получаем файл из формы
        file, header, err := r.FormFile("file")
        if err != nil {
            http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
            return
        }
        defer file.Close()

        // Прочитать данные из файла
        content, err = io.ReadAll(file)
        if err != nil {
            http.Error(w, "Failed to read file", http.StatusInternalServerError)
            return
        }
        
        filename = header.Filename
    } else {
        // Если не multipart (тесты отправляют текст напрямую)
        // Читаем тело запроса напрямую
        content, err = io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Failed to read request body", http.StatusInternalServerError)
            return
        }
        defer r.Body.Close()
        
        // Для тестов используем дефолтное имя файла
        filename = "input.txt"
    }

    // Передать эти данные в функцию автоопределения
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }

    // Создать локальный файл
    timestamp := time.Now().UTC().String()
    ext := filepath.Ext(filename)
    if ext == "" {
        ext = ".txt"
    }
    
    // Очищаем timestamp для имени файла
    safeTimestamp := strings.ReplaceAll(strings.ReplaceAll(timestamp, " ", "_"), ":", "-")
    outputFilename := "converted_" + safeTimestamp + ext

    // Записать в локальный файл результат конвертации
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

    // Вернуть результат конвертации строки
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(converted))
}