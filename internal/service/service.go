package service

import (
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Определяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }
    
    // Ищем первую не-пробельную точку или тире
    for i := 0; i < len(trimmed); i++ {
        if trimmed[i] == '.' || trimmed[i] == '-' {
            return true
        }
        if trimmed[i] != ' ' && trimmed[i] != '\t' && trimmed[i] != '\n' {
            return false
        }
    }
    
    return false
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    // Убираем BOM если есть (иногда бывает в начале файлов)
    input = strings.TrimPrefix(input, "\ufeff")
    
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