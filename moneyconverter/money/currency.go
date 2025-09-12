package money

import (
	"strings"
	"unicode"
)

type Currency struct {
	isoCode   string
	precision byte
}

var (
	CurrencyParseInvalidCodeError = Error("The given code does not match an ISO 4217 currency code.")
)

func ParseCurrency(input string) (Currency, error) {

	if validateCurrencyInput(input) != nil {
		// ???
	}

	return Currency{isoCode: strings.ToUpper(input)}, nil
}

func validateCurrencyInput(input string) error {
	if len(input) != 3 {
		return CurrencyParseInvalidCodeError
	}

	// Go has no native way of checking that a letter is english,
	// But we know that ASCII is 127 chars.
	// If a rune (uint32) is > 127 then its not an ASCII char - we can then check it's a letter.
	// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters

	for _, c := range input {
		if c > unicode.MaxASCII {
			if !unicode.IsLetter(c) {
				return CurrencyParseInvalidCodeError
			}
		}
	}

	return nil
}
