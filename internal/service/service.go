package service

import (
    "strings"
)

// Создаем свою карту преобразования Морзе-Текст
var morseToTextMap = map[string]string{
    ".-":     "А",     "-...":   "Б",     ".--":    "В",     "--.":    "Г",
    "-..":    "Д",     ".":      "Е",     "...-":   "Ж",     "--..":   "З",
    "..":     "И",     ".---":   "Й",     "-.-":    "К",     ".-..":   "Л",
    "--":     "М",     "-.":     "Н",     "---":    "О",     ".--.":   "П",  
    ".-.":    "Р",     "...":    "С",     "-":      "Т",     "..-":    "У",
    "..-.":   "Ф",     "....":   "Х",     "-.-.":   "Ц",     "---.":   "Ч",
    "----":   "Ш",     "--.-":   "Щ",     "--.--":  "Ъ",     "-.--":   "Ы",
    "-..-":   "Ь",     "..-..":  "Э",     "..--":   "Ю",     ".-.-":   "Я",
    ".----":  "1",     "..---":  "2",     "...--":  "3",     "....-":  "4",
    ".....":  "5",     "-....":  "6",     "--...":  "7",     "---..":  "8",
    "----.":  "9",     "-----":  "0",     "/":      " ",     " ":      "",
}


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
        
        return convertMorseToText(trimmed), nil
    }
    
    
    return convertTextToMorse(trimmed), nil
}

func convertMorseToText(morse string) string {
    words := strings.Split(morse, " / ")
    var result []string
    
    for _, word := range words {
        letters := strings.Split(strings.TrimSpace(word), " ")
        var wordText strings.Builder
        
        for _, letter := range letters {
            if text, ok := morseToTextMap[letter]; ok {
                wordText.WriteString(text)
            }
        }
        
        if wordText.Len() > 0 {
            result = append(result, wordText.String())
        }
    }
    
    return strings.Join(result, " ")
}

func convertTextToMorse(text string) string {
    text = strings.ToUpper(text)
    var result []string
    
    for _, r := range text {
        ch := string(r)
        
        morse := findMorseForRune(r)
        if morse != "" {
            result = append(result, morse)
        } else if ch == " " {
            result = append(result, "/")
        }
    }
    
    return strings.Join(result, " ")
}

func findMorseForRune(r rune) string {
    for morse, text := range morseToTextMap {
        if text == string(r) {
            return morse
        }
    }
    return ""
}