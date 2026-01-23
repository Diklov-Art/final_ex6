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

    // Используем ContainsFunc как указано в ТЗ
    // Проверяем, что строка содержит ТОЛЬКО допустимые символы Морзе
    return !strings.ContainsFunc(trimmed, func(r rune) bool {
        // Допустимые символы: точка, тире, пробел, слэш
        return !(r == '.' || r == '-' || r == ' ' || r == '/')
    })
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    if isMorseCode(input) {
        // Конвертируем код Морзе в текст
        result := morse.ToText(input)
        if result == "" {
            return "", nil
        }
        return result, nil
    } else {
        // Конвертируем текст в код Морзе
        result := morse.ToMorse(input)
        if result == "" {
            return "", nil
        }
        return result, nil
    }
}