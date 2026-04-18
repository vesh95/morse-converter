package service

import (
	"testing"
)

func TestIsMorse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Пустые и невалидные строки
		{"empty string", "", false},
		{"only letters", "abc", false},
		{"letters and morse", ".-a", false},
		{"numbers", "123", false},
		{"mixed invalid", ".- 123 abc", false},

		// Валидный код Морзе
		{"single dot", ".", true},
		{"single dash", "-", true},
		{"morse A", ".-", true},
		{"morse B", "-...", true},
		{"morse with space", ".- -...", true},
		{"multiple spaces", ".-   -...", true},
		{"with tabs", "\t.-", true},
		{"with newlines", "\n.-", true},
		{"surrounding spaces", " .- ", true},
		{"question mark", "..--..", true},
		{"period", "......", true},
		{"hello", ".... . .-.. .-.. ---", true},

		// Граничные случаи
		{"only spaces", "   ", false},
		{"only tabs", "\t\t\t", false},
		{"spaces and tabs", " \t \n ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMorse(tt.input)
			if result != tt.expected {
				t.Errorf("isMorse(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"text to morse", "АБ", ".- -..."},
		{"morse to text", ".- -...", "АБ"},
		{"hello text", "ПРИВЕТ", ".--. .-. .. .-- . -"},
		{"hello morse", ".--. .-. .. .-- . -", "ПРИВЕТ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Сonvert(tt.input)
			if result != tt.expected {
				t.Errorf("Сonvert(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
