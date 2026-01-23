package service

import (
    "fmt"
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var debug = false // можно включить для отладки

// Определяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }
    
    if debug {
        fmt.Printf("DEBUG isMorseCode: %q\n", trimmed)
        if len(trimmed) > 100 {
            fmt.Printf("First 100 chars: %q\n", trimmed[:100])
        }
    }
    
    // Быстрая проверка: если начинается с точки или тире - вероятно код Морзе
    for _, r := range trimmed {
        if r == ' ' || r == '\t' || r == '\n' {
            continue
        }
        if r == '.' || r == '-' {
            if debug {
                fmt.Printf("DEBUG: First non-space char is '.' or '-', returning true\n")
            }
            return true
        }
        break
    }
    
    return false
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    if debug {
        fmt.Printf("DEBUG Convert input: %q\n", input)
    }
    
    if isMorseCode(input) {
        // Конвертируем код Морзе в текст
        result := morse.ToText(input)
        if debug {
            fmt.Printf("DEBUG: Treated as morse, result: %q\n", result)
        }
        return result, nil
    } else {
        // Конвертируем текст в код Морзе
        result := morse.ToMorse(input)
        if debug {
            fmt.Printf("DEBUG: Treated as text, result: %q\n", result)
        }
        return result, nil
    }
}