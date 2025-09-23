package money

import (
	"errors"
	"strings"
	"unicode"
)

// List of currencies in currencies.json obtained from here : https://github.com/ourworldincode/currency/blob/main/currencies.json

var (
	errCurrencyDataNotLoaded    = errors.New("No currencies found")
	errCurrencyListEmpty        = errors.New("The provided currency data is empty")
	errCurrencyValidationError  = errors.New("Currency object is invalid")
	errCurrencyDataRead         = errors.New("Could not get list of currencies")
	errCurrencyNotFound         = errors.New("The given input could not be found")
	errCurrencyInputValidation  = errors.New("The provided input was not valid")
	errCurrencyParseInvalidCode = errors.New("The given code does not match an ISO 4217 currency code.")
)

var (
	CurrencyDataNotLoadedError    = errCurrencyDataNotLoaded
	CurrencyDataReadError         = errCurrencyDataRead
	CurrencyNotFoundError         = errCurrencyNotFound
	CurrencyInputValidationError  = errCurrencyInputValidation
	CurrencyListEmptyError        = errCurrencyListEmpty
	CurrencyParseInvalidCodeError = errCurrencyParseInvalidCode
	CurrencyValidationError       = errCurrencyValidationError
)

type Currency struct {
	ISOCode   string
	Precision byte
}

func (c *Currency) validate() error {
	valid := len(c.ISOCode) == 3 && c.Precision != 0
	if !valid {
		return errCurrencyValidationError
	}
	return nil
}

type CurrencyList = map[string]Currency

type CurrencyParser struct {
	currencies CurrencyList
}

func NewCurrencyParser(data CurrencyData) (*CurrencyParser, error) {
	currencies, err := data.GetCurrencies()

	if err != nil {
		return &CurrencyParser{}, errCurrencyDataNotLoaded
	}

	return &CurrencyParser{currencies: currencies}, nil

}

func (parser *CurrencyParser) ParseCurrency(input string) (Currency, error) {

	err := validateCurrencyInput(input)

	if err != nil {
		return Currency{}, err
	}

	currency, exists := parser.currencies[strings.ToUpper(input)]

	if !exists {
		return Currency{}, errCurrencyNotFound
	}

	return currency, nil
}

func validateCurrencyInput(input string) error {
	if len(input) != 3 {
		return errCurrencyInputValidation
	}

	// Go has no native way of checking that a letter is english,
	// But we know that ASCII is 127 chars.
	// If a rune (uint32) is > 127 then its not an ASCII char - we can fail validation.
	// If it is, then we can check that the character is a letter.
	// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters

	for _, c := range input {
		if c < unicode.MaxASCII {
			if !unicode.IsLetter(c) {
				return errCurrencyInputValidation
			}
		} else {
			// Fail for non ascii characters.
			return errCurrencyInputValidation
		}
	}

	return nil
}
