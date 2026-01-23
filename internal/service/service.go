package service

import (
    "strings"
)

// Полная собственная реализация конвертации Морзе

var morseAlphabet = map[string]string{
    "А": ".-",   "Б": "-...", "В": ".--",  "Г": "--.",  "Д": "-..",
    "Е": ".",    "Ж": "...-", "З": "--..", "И": "..",   "Й": ".---",
    "К": "-.-",  "Л": ".-..", "М": "--",   "Н": "-.",   "О": "---",
    "П": ".--.", "Р": ".-.",  "С": "...",  "Т": "-",    "У": "..-",
    "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.", "Ш": "----",
    "Щ": "--.-", "Ъ": "--.--","Ы": "-.--", "Ь": "-..-", "Э": "..-..",
    "Ю": "..--", "Я": ".-.-", " ": "/",
    
    "1": ".----", "2": "..---", "3": "...--", "4": "....-", "5": ".....",
    "6": "-....", "7": "--...", "8": "---..", "9": "----.", "0": "-----",
}

var textAlphabet = map[string]string{}

func init() {
    // Создаем обратное отображение
    for k, v := range morseAlphabet {
        textAlphabet[v] = k
    }
}

// Определяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
    trimmed := strings.TrimSpace(s)
    if trimmed == "" {
        return false
    }

    // Используем strings.ContainsFunc как указано в ТЗ
    if strings.ContainsFunc(trimmed, func(r rune) bool {
        return !(r == '.' || r == '-' || r == ' ' || r == '/' || r == '\t' || r == '\n')
    }) {
        return false
    }
    
    // Должна быть хотя бы одна точка или тире
    for _, r := range trimmed {
        if r == '.' || r == '-' {
            return true
        }
    }
    
    return false
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    // Убираем BOM если есть
    input = strings.TrimPrefix(input, "\ufeff")
    input = strings.TrimSpace(input)
    
    if isMorseCode(input) {
        // Конвертируем код Морзе в текст
        return convertMorseToText(input), nil
    } else {
        // Конвертируем текст в код Морзе
        return convertTextToMorse(input), nil
    }
}

func convertMorseToText(morse string) string {
    words := strings.Split(morse, " / ")
    var result []string
    
    for _, word := range words {
        letters := strings.Split(strings.TrimSpace(word), " ")
        var decodedWord strings.Builder
        
        for _, letter := range letters {
            if text, ok := textAlphabet[letter]; ok {
                decodedWord.WriteString(text)
            }
        }
        
        if decodedWord.Len() > 0 {
            result = append(result, decodedWord.String())
        }
    }
    
    return strings.Join(result, " ")
}

func convertTextToMorse(text string) string {
    text = strings.ToUpper(text)
    var result []string
    
    for _, r := range text {
        ch := string(r)
        if morse, ok := morseAlphabet[ch]; ok {
            result = append(result, morse)
        } else if ch == " " {
            result = append(result, "/")
        }
    }
    
    return strings.Join(result, " ")
}