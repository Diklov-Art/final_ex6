package service

import (
    "fmt"
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    trimmed := strings.TrimSpace(input)
    
    // Логирование для отладки
    fmt.Printf("SERVICE DEBUG: Input: %q\n", trimmed)
    if len(trimmed) > 50 {
        fmt.Printf("SERVICE DEBUG: First 50 chars: %q\n", trimmed[:50])
    }
    
    // Проверяем, является ли строка кодом Морзе
    isMorse := true
    hasDotsOrDashes := false
    
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            hasDotsOrDashes = true
        } else if !(r == ' ' || r == '/' || r == '\t' || r == '\n') {
            isMorse = false
            break
        }
    }
    
    fmt.Printf("SERVICE DEBUG: isMorse: %v, hasDotsOrDashes: %v\n", isMorse, hasDotsOrDashes)
    
    if isMorse && hasDotsOrDashes {
        // Морзе -> Текст
        result := morse.ToText(trimmed)
        fmt.Printf("SERVICE DEBUG: morse.ToText result: %q\n", result)
        
        // Проверяем, что получили
        if strings.Contains(trimmed, ".--.") {
            fmt.Printf("SERVICE DEBUG: Input contains '.--.' (code for П)\n")
        }
        
        return result, nil
    }
    
    
    result := morse.ToMorse(trimmed)
    fmt.Printf("SERVICE DEBUG: morse.ToMorse result: %q\n", result)
    return result, nil
}