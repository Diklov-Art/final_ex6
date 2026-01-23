package service

import (
    "strings"
    "unicode"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Определяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }

    // Простая проверка: если больше 50% символов - точки или тире
    totalChars := len(trimmed)
    dotsAndDashes := 0
    
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            dotsAndDashes++
        }
    }
    
    // Если более 30% символов - точки или тире, считаем что это код Морзе
    if totalChars > 0 && float64(dotsAndDashes)/float64(totalChars) > 0.3 {
        return true
    }
    
    return false
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    if isMorseCode(input) {
        // Конвертируем код Морзе в текст
        result := morse.ToText(input)
        return result, nil
    } else {
        // Конвертируем текст в код Морзе
        result := morse.ToMorse(input)
        return result, nil
    }
}