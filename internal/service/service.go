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
        if result == "" {
            return "", nil
        }
        return result, nil
    } else {
        
        result := morse.ToMorse(input)
        if result == "" {
            return "", nil
        }
        return result, nil
    }
}


func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }

    
    if strings.ContainsFunc(trimmed, func(r rune) bool {
        // Допустимые символы в коде Морзе: точка, тире, пробел, слэш
        return !(r == '.' || r == '-' || r == ' ' || r == '/' || r == '\t' || r == '\n')
    }) {
        return false
    }
    
   
    hasDotsOrDashes := false
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            hasDotsOrDashes = true
            break
        }
    }
    
    return hasDotsOrDashes
}