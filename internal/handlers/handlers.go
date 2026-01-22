package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	// Отправляем HTML файл
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму с файлом
	err := r.ParseMultipartForm(10 << 20) // 10 MB максимум
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

	// Читаем содержимое файла
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Конвертируем содержимое
	content := string(fileBytes)
	converted, err := service.Convert(content)
	if err != nil {
		http.Error(w, "Failed to convert content", http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл для результата
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}

	outputFilename := fmt.Sprintf("converted_%s%s", timestamp, ext)

	// Создаем файл для записи результата
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Отправляем результат и информацию о файле
	response := fmt.Sprintf("Conversion completed!\n\nResult saved to: %s\n\nConverted content:\n%s",
		outputFilename, converted)

	_, err = w.Write([]byte(response))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
