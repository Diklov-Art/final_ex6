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
    
    if isMorse && hasDotsOrDashes {
        
        return morse.ToText(trimmed), nil
    }
    
    
    return morse.ToMorse(input), nil
}