package valueobject

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestCPF(t *testing.T) {
	suite.Run(t, new(testCPFSuite))
}

type testCPFSuite struct {
	suite.Suite
}

func (s *testCPFSuite) TestRemoveNonDigits() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"should remove non digits": {
			input:    "12345Dl@p",
			expected: "12345",
		},
		"should dont remove digits": {
			input:    "123456",
			expected: "123456",
		},
		"should not fail on empty input": {
			input:    "",
			expected: "",
		},
		"should remove empty spaces": {
			input:    "1234 56  78",
			expected: "12345678",
		},
	}
	for name, tt := range tests {
		s.T().Run(name, func(t *testing.T) {
			result := removeNonDigits(tt.input)
			if result != tt.expected {
				t.Fatalf("Result is %s not equal to %s", result, tt.expected)
			}
		})
	}
}

func (s *testCPFSuite) TestAllDigitsEqual() {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"should be false all equal 1": {
			input:    "11111111111",
			expected: false,
		},
		"should be false all equal 2": {
			input:    "22222222222",
			expected: false,
		},
		"should be ok on different values": {
			input:    "11144477705",
			expected: true,
		},
	}
	for name, tt := range tests {
		s.T().Run(name, func(t *testing.T) {
			result := allDigitsEqual(tt.input)
			if result != tt.expected {
				t.Fatalf("Result is %s not valid", tt.input)
			}
		})
	}
}

func (s *testCPFSuite) TestCalculateDigit() {
	tests := map[string]struct {
		input    string
		factor   int
		expected string
	}{
		"should be remainder of 3 on factor first digit": {
			input:    "11144477705",
			factor:   factor_first_digit,
			expected: "3",
		},
		"should be remainder of 5 on factor second digit": {
			input:    "11144477705",
			factor:   factor_second_digit,
			expected: "0",
		},
	}
	for name, tt := range tests {
		s.T().Run(name, func(t *testing.T) {
			result := calculateDigit(tt.input, tt.factor)
			if result != tt.expected {
				t.Fatalf("Result is %s not valid, expected %s", tt.expected, result)
			}
		})
	}
}

func (s *testCPFSuite) TestValidateCPF() {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"should be valid CPF value 97456321558": {
			input:    "97456321558",
			expected: true,
		},
		"should be valid CPF value 71428793860": {
			input:    "71428793860",
			expected: true,
		},
		"should be valid CPF value 87748248800": {
			input:    "87748248800",
			expected: true,
		},
		"should be invalid CPF value 11144477705": {
			input:    "11144477705",
			expected: false,
		},
		"should be invalid CPF value 11111111111": {
			input:    "11111111111",
			expected: false,
		},
		"should be invalid CPF 123": {
			input:    "123",
			expected: false,
		},
		"should be invalid CPF 1234566789123456789": {
			input:    "1234566789123456789",
			expected: false,
		},
		"should be invalid empty CPF": {
			input:    "",
			expected: false,
		},
	}
	for name, tt := range tests {
		s.T().Run(name, func(t *testing.T) {
			result := validate(tt.input)
			if result != tt.expected {
				t.Fatalf("Result is %s not valid", tt.input)
			}
		})
	}
}
