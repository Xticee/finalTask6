package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Decode(str string) (string, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return "", errors.New("empty string")
	}

	if isMorseCode(str) {
		return morse.ToText(str), nil
	}

	return morse.ToMorse(str), nil
}

func isMorseCode(input string) bool {

	cleaned := strings.ReplaceAll(input, " ", "")

	for _, rune := range cleaned {
		if rune != '.' && rune != '-' {
			return false
		}
	}

	return true
}
