package testutils

import (
	"learngo-pockets/moneyconverter/money"
	"testing"
)

type testCurrencyData struct {
	input money.CurrencyList
}

func (tcd *testCurrencyData) GetCurrencies() (money.CurrencyList, error) {
	return tcd.input, nil
}

var mockCurrencyData = map[string]money.Currency{
	"USD": {ISOCode: "USD", Precision: 2},
	"GBP": {ISOCode: "GBP", Precision: 2},
	"JPY": {ISOCode: "JPY", Precision: 0},
}

func SetupTestCurrencyParser(t *testing.T) *money.CurrencyParser {
	t.Helper()

	parser, err := money.NewCurrencyParser(&testCurrencyData{input: mockCurrencyData})

	if err != nil {
		t.Fatalf(err.Error())
	}

	return parser
}
func SetupTestConversionData(t *testing.T, amtDecimal string, fromCurrencyCode string, toCurrencyCode string) (money.Amount, money.Currency) {
	t.Helper()

	parser := SetupTestCurrencyParser(t)

	// Build a decimal.
	fromAmount, err := money.ParseDecimal(amtDecimal)

	if err != nil {
		t.Fatalf(err.Error())
	}

	// Build currencies
	fromCurrency, err := parser.ParseCurrency(fromCurrencyCode)

	if err != nil {
		t.Fatalf(err.Error())
	}

	toCurrency, err := parser.ParseCurrency(toCurrencyCode)

	if err != nil {
		t.Fatalf(err.Error())
	}

	// Build amount with initial currency
	amt, err := money.NewAmount(fromAmount, fromCurrency)

	if err != nil {
		t.Fatalf(err.Error())
	}

	// Return the amount, and the destination currency for use in conversion tests.
	return amt, toCurrency
}
