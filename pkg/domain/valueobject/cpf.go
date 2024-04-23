package valueobject

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCPF = errors.New("invalid CPF")
)

type Cpf string

func CpfFromString(s string) (Cpf, error) {
	if !validate(s) {
		return "", ErrInvalidCPF
	}
	return Cpf(s), nil
}

func (c Cpf) String() string {
	return string(c)
}

const (
	cpf_valid_length    = 11
	factor_first_digit  = 10
	factor_second_digit = 11
)

func validate(rawCPF string) bool {
	if rawCPF == "" {
		return false
	}
	cpf := removeNonDigits(rawCPF)
	if !isValidLength(cpf) {
		return false
	}
	if !allDigitsEqual(cpf) {
		return false
	}
	firstDigit := calculateDigit(cpf, factor_first_digit)
	secondDigit := calculateDigit(cpf, factor_second_digit)
	digitExtracted := string(cpf[9:])
	return digitExtracted == firstDigit+secondDigit
}

func removeNonDigits(rawCPF string) string {
	expr := regexp.MustCompile(`\D`)
	ret := string(expr.ReplaceAll([]byte(rawCPF), []byte("")))
	return ret
}

func isValidLength(cpf string) bool {
	return len(cpf) == cpf_valid_length
}

func allDigitsEqual(cpf string) bool {
	firstDigits := cpf[0]
	for _, digit := range cpf {
		if digit != rune(firstDigits) {
			return true
		}
	}
	return false
}

func calculateDigit(cpf string, factor int) string {
	total := 0
	for _, digit := range cpf {
		if factor > 1 {
			numDigit, _ := strconv.Atoi(string(digit))
			total += numDigit * factor
			factor = factor - 1
		}
	}
	remainder := total % 11
	if remainder < 2 {
		return "0"
	}
	return strconv.Itoa(11 - remainder)
}
