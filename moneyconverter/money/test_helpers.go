package money

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type testCurrencyData struct {
	input CurrencyList
}

func (tcd *testCurrencyData) GetCurrencies() (CurrencyList, error) {
	return tcd.input, nil
}

var mockCurrencyData = map[string]Currency{
	"USD": {ISOCode: "USD", Precision: 2},
	"GBP": {ISOCode: "GBP", Precision: 2},
	"JPY": {ISOCode: "JPY", Precision: 0},
}

func SetupTestCurrencyParser(t *testing.T) *CurrencyParser {
	t.Helper()

	parser, err := NewCurrencyParser(&testCurrencyData{input: mockCurrencyData})

	if err != nil {
		t.Fatal(err.Error())
	}

	return parser
}

func SetupTestConversionData(t *testing.T, amtDecimal string, fromCurrencyCode string, toCurrencyCode string) (Amount, Currency) {
	t.Helper()

	parser := SetupTestCurrencyParser(t)

	// Build a decimal.
	fromAmount, err := ParseDecimal(amtDecimal)

	if err != nil {
		t.Fatal(err.Error())
	}

	// Build currencies
	fromCurrency, err := parser.ParseCurrency(fromCurrencyCode)

	if err != nil {
		t.Fatal(err.Error())
	}

	toCurrency, err := parser.ParseCurrency(toCurrencyCode)

	if err != nil {
		t.Fatal(err.Error())
	}

	// Build amount with initial currency
	amt, err := NewAmount(fromAmount, fromCurrency)

	if err != nil {
		t.Fatal(err.Error())
	}

	// Return the amount, and the destination currency for use in conversion tests.
	return amt, toCurrency
}

type mockRateService struct {
	rate ExchangeRate
}

func (mrs *mockRateService) getRate(string) (ExchangeRate, error) {
	return mrs.rate, nil
}

func SetupTestConverter(t *testing.T, rateString string) *converter {
	spew.Dump("dumping rateStr")
	spew.Dump(rateString)
	rate, err := ParseDecimal(rateString)

	if err != nil {
		t.Fatal(err.Error())
	}

	converter := NewConverter(&mockRateService{rate: rate})
	return converter
}
