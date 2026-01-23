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

    // Используем ContainsFunc как указано 
    // Если строка содержит ЛЮБОЙ символ, кроме допустимых для кода Морзе
    // то это не код Морзе
    if strings.ContainsFunc(trimmed, func(r rune) bool {
        // Допустимые символы в коде Морзе
        return !(r == '.' || r == '-' || r == ' ' || r == '/')
    }) {
        return false
    }
    
    // Дополнительная проверка: строка должна содержать хотя бы одну точку или тире
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            return true
        }
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