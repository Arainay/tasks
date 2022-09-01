package lib

import (
	"regexp"
	"strings"
)

var logins = map[string]string{}

// На вход подаются n логинов, для каждого логина вам нужно сказать, можно ли его зарегистрировать. Логин должен соответствовать следующим правилам:
// Логин — это последовательность из латинских букв в любом регистре, цифр и символов «_» или «-» (подчеркивание и дефис).
// Логин не должен начинаться с дефиса.
// Логин должен иметь длину от 2 до 24 символов.
// Логины, которые отличаются только регистром, считаются одинаковыми.
func Register(login string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{1,23}$`, login)

	if !matched {
		return false
	}

	key := strings.ToLower(login)
	_, isExists := logins[key]

	if isExists {
		return false
	}

	logins[key] = login

	return true
}
