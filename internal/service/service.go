package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/morse-converter/pkg/morse"
)

// Определяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
	// Код Морзе содержит только точки, тире, пробелы и символы / для разделения слов
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return false
	}

	// Проверяем, содержит ли строка только допустимые символы Морзе
	for _, r := range trimmed {
		if !(r == '.' || r == '-' || r == ' ' || r == '/') {
			return false
		}
	}

	return true
}

// Convert автоматически определяет тип строки и конвертирует ее
func Convert(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	if isMorseCode(input) {
		// Конвертируем код Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Конвертируем текст в код Морзе
		return morse.ToMorse(input), nil
	}
}
