package handlers

import (
	"io"
	"net/http"
	"os"
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

	var content []byte
	var err error

	// Пробуем получить файл через multipart/form-data
	if r.Header.Get("Content-Type") == "multipart/form-data" {
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
	} else {
		// Если не multipart, читаем тело запроса напрямую
		content, err = io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
	}

	// Конвертируем содержимое
	converted, err := service.Convert(string(content))
	if err != nil {
		http.Error(w, "Failed to convert content", http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл для результата
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	outputFilename := "converted_" + timestamp + ".txt"

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

	// Только результат конвертации (без дополнительного текста)
	_, err = w.Write([]byte(converted))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
