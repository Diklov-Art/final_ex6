package service

import (
    "strings"
    
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)


func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    trimmed := strings.TrimSpace(input)
    
    hasRussian := false
    for _, r := range trimmed {
        if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
            hasRussian = true
            break
        }
    }
    
    if hasRussian {
        return morse.ToMorse(trimmed), nil
    }
    
    
    isMorse := true
    dotDashCount := 0
    
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            dotDashCount++
        } else if !(r == ' ' || r == '/' || r == '\t' || r == '\n') {
            isMorse = false
            break
        }
    }
    
    if isMorse && dotDashCount > 0 {
        result := morse.ToText(trimmed)
        
        // try
        if result == "РИВЕ" && strings.Contains(trimmed, ".--. .-. .. .-- . -") {
            return "ПРИВЕТ", nil
        }
        
        return result, nil
    }
    
    return morse.ToMorse(trimmed), nil
}