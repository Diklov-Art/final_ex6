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

/
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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var content []byte
    var filename string
    var err error

    
    contentType := r.Header.Get("Content-Type")
    
    if strings.Contains(contentType, "multipart/form-data") {
       
        err = r.ParseMultipartForm(10 << 20) 
        if err != nil {
            http.Error(w, "Failed to parse form", http.StatusInternalServerError)
            return
        }

       
        file, header, err := r.FormFile("file")
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
        
        filename = header.Filename
    } else {
        
        content, err = io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Failed to read request body", http.StatusInternalServerError)
            return
        }
        defer r.Body.Close()
        
        filename = "input.txt" 
    }

    
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }

    
    timestamp := time.Now().UTC().String()
    
    
    ext := filepath.Ext(filename)
    if ext == "" {
        ext = ".txt"
    }
    
    
    safeTimestamp := strings.ReplaceAll(strings.ReplaceAll(
        strings.ReplaceAll(timestamp, " ", "_"), 
        ":", "-"), 
        "+", "_")
    
    outputFilename := "converted_" + safeTimestamp + ext

    
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

    
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(converted))
}