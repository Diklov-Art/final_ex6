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
    
    // Проверяем наличие русских букв
    for _, r := range trimmed {
        if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
            return false // Русские буквы - это текст
        }
    }
    
    // Проверяем наличие английских букв (латиницы)
    for _, r := range trimmed {
        if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
            return false // Латинские буквы - это текст
        }
    }
    
    // Если дошли сюда, проверяем на код Морзе
    // Строка должна содержать точки или тире
    hasDotsOrDashes := false
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            hasDotsOrDashes = true
            break
        }
    }
    
    return hasDotsOrDashes
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