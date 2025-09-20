package money

import (
	"strings"
	"unicode"
)

// List of currencies in currencies.json obtained from here : https://github.com/ourworldincode/currency/blob/main/currencies.json

var (
	CurrencyDataNotLoadedError    = Error("No currencies found")
	CurrencyDataReadError         = Error("Could not get list of currencies")
	CurrencyNotFoundError         = Error("The given input could not be found")
	CurrencyInputValidationError  = Error("The provided input was not valid")
	CurrencyParseInvalidCodeError = Error("The given code does not match an ISO 4217 currency code.")
)

type Currency struct {
	isoCode   string
	precision byte
}

type CurrencyList = map[string]Currency

type CurrencyParser struct {
	currencies CurrencyList
}

func NewCurrencyParser(data CurrencyData) (*CurrencyParser, Error) {
	currencies, err := data.getCurrencies()

	if err != nil {
		return &CurrencyParser{}, CurrencyDataReadError
	}

	return &CurrencyParser{currencies: currencies}, ""

}

func (parser *CurrencyParser) ParseCurrency(input string) (Currency, Error) {

	err := validateCurrencyInput(input)

	if err != nil {
		return Currency{}, CurrencyDataReadError
	}

	_, exists := parser.currencies[strings.ToUpper(input)]

	if !exists {
		return Currency{}, CurrencyInputValidationError
	}

	return Currency{isoCode: strings.ToUpper(input)}, ""
}

func validateCurrencyInput(input string) error {
	if len(input) != 3 {
		return Error(CurrencyParseInvalidCodeError)
	}

	// Go has no native way of checking that a letter is english,
	// But we know that ASCII is 127 chars.
	// If a rune (uint32) is > 127 then its not an ASCII char - we can then check it's a letter.
	// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters

	for _, c := range input {
		if c > unicode.MaxASCII {
			if !unicode.IsLetter(c) {
				return Error(CurrencyParseInvalidCodeError)
			}
		}
	}

	return nil
}
