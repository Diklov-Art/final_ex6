package handlers

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)


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

    
    err := r.ParseMultipartForm(10 << 20) 
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusInternalServerError)
        return
    }

   
    file, header, err := r.FormFile("myFile") 
    if err != nil {
        http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    
    content, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Failed to read file", http.StatusInternalServerError)
        return
    }

    // DEBUG
    fmt.Printf("HANDLER DEBUG: Filename: %s\n", header.Filename)
    fmt.Printf("HANDLER DEBUG: Content length: %d bytes\n", len(content))
    if len(content) > 0 {
        fmt.Printf("HANDLER DEBUG: First 100 chars of content: %q\n", string(content[:min(100, len(content))]))
    }

    
    converted, err := service.Convert(string(content))
    if err != nil {
        http.Error(w, "Failed to convert content", http.StatusInternalServerError)
        return
    }

    timestamp := time.Now().UTC().String()
    
    
    ext := filepath.Ext(header.Filename)
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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}