package service

import (
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)


func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    // Определяем, является ли строка кодом Морзе
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


func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }

    
    
    // Считаем количество точек и тире
    dotsAndDashes := 0
    otherAllowedChars := 0 // пробелы, слэши
    
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            dotsAndDashes++
        } else if r == ' ' || r == '/' || r == '\t' || r == '\n' {
            otherAllowedChars++
        } else {
            // Нашли символ, который недопустим в коде Морзе
            return false
        }
    }
    
    
    return dotsAndDashes > 0
}