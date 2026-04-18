package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Сonvert(input string) string {
	var converted string
	converter := morse.NewConverter(morse.DefaultMorse, morse.WithLowercaseHandling(true))
	if isMorse(input) {
		converted = converter.ToText(input)
	} else {
		converted = converter.ToMorse(input)
	}

	return converted
}

func isMorse(input string) bool {
	// Пустая строка не является кодом Морзе
	if len(input) == 0 {
		return false
	}

	// Проверяем, содержит ли строка хотя бы один символ Морзе (точку или тире)
	if !strings.ContainsAny(input, ".-") {
		return false
	}

	// Удаляем все допустимые символы Морзе и пробельные символы
	// Если после удаления останутся символы - это не код Морзе
	cleaned := strings.Map(func(r rune) rune {
		switch r {
		case '.', '-', ' ', '\t', '\n', '\r':
			return -1
		default:
			return r
		}
	}, input)

	// Если после очистки остались символы - это не код Морзе
	if len(cleaned) > 0 {
		return false
	}

	return true
}
