package service

import (
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    // Очищаем входные данные
    input = strings.TrimSpace(input)
    
    // Очень простая эвристика:
    // Если строка содержит русские буквы - это текст
    hasRussian := false
    for _, r := range input {
        if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
            hasRussian = true
            break
        }
    }
    
    if hasRussian {
        // Текст -> Морзе
        return morse.ToMorse(input), nil
    }
    
    // Иначе пробуем как код Морзе
    result := morse.ToText(input)
    if result == "" {
        // Если не получилось декодировать, возможно это английский текст
        return morse.ToMorse(input), nil
    }
    
    return result, nil
}