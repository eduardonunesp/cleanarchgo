package valueobject

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCPF        = errors.New("invalid CPF")
	removeNonDigitsRegex = regexp.MustCompile(`\D`)
)

type Cpf struct {
	value string
}

func CpfFromString(s string) (Cpf, error) {
	var cpf Cpf
	if !validate(s) {
		return cpf, ErrInvalidCPF
	}
	cpf.value = s
	return cpf, nil
}

func (c Cpf) String() string {
	return c.value
}

const (
	CpfValidLength    = 11
	FactorFirstDigit  = 10
	FactorSecondDigit = 11
	RemainderValue    = 2
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
	firstDigit := calculateDigit(cpf, FactorFirstDigit)
	secondDigit := calculateDigit(cpf, FactorSecondDigit)
	digitExtracted := cpf[9:]
	return digitExtracted == firstDigit+secondDigit
}

func removeNonDigits(rawCPF string) string {
	ret := string(removeNonDigitsRegex.ReplaceAll([]byte(rawCPF), []byte("")))
	return ret
}

func isValidLength(cpf string) bool {
	return len(cpf) == CpfValidLength
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
	remainder := total % CpfValidLength
	if remainder < RemainderValue {
		return "0"
	}
	return strconv.Itoa(CpfValidLength - remainder)
}
