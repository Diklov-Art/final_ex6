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

    // Используем strings.ContainsFunc как указано в ТЗ
    // Проверяем, содержит ли строка символы, которые НЕ являются допустимыми для кода Морзе
    if strings.ContainsFunc(trimmed, func(r rune) bool {
        // Допустимые символы в коде Морзе: точка, тире, пробел, слэш
        return !(r == '.' || r == '-' || r == ' ' || r == '/' || r == '\t' || r == '\n')
    }) {
        return false
    }
    
    // Также проверяем, что строка содержит хотя бы одну точку или тире
    // (иначе строка из одних пробелов будет считаться кодом Морзе)
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