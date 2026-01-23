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

    // Проверяем, содержит ли строка русские буквы
    hasCyrillic := false
    for _, r := range trimmed {
        if unicode.Is(unicode.Cyrillic, r) {
            hasCyrillic = true
            break
        }
    }
    
    // Если есть русские буквы - это текст, не код Морзе
    if hasCyrillic {
        return false
    }
    
    // Проверяем, содержит ли строка только допустимые символы Морзе
    // Код Морзе может содержать: точку, тире, пробел, слэш
    // Также могут быть цифры, которые кодируются точками и тире
    
    // Считаем количество точек и тире
    dotsAndDashes := 0
    otherChars := 0
    
    for _, r := range trimmed {
        switch r {
        case '.', '-':
            dotsAndDashes++
        case ' ', '/', '\t', '\n':
            // Пробельные символы разрешены
        default:
            otherChars++
        }
    }
    
    // Если есть другие символы (кроме точек, тире и пробелов) - это не код Морзе
    if otherChars > 0 {
        return false
    }
    
    // Если есть хотя бы одна точка или тире - считаем что это код Морзе
    return dotsAndDashes > 0
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