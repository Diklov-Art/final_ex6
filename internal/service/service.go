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
    
    
    asText := morse.ToText(trimmed)
    asMorse := morse.ToMorse(trimmed)
    
    
    if asText != "" && asText != trimmed {
        hasRussian := false
        for _, r := range asText {
            if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') {
                hasRussian = true
                break
            }
        }
        
        if hasRussian {
            return asText, nil
        }
    }
    
    
    return asMorse, nil
}