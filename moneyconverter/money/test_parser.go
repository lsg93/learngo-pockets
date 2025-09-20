package money

type TestCurrencyData struct {
	input CurrencyList
}

func (tcd *TestCurrencyData) getCurrencies() (map[string]Currency, error) {
	return tcd.input, nil
}

func setupTestCurrencyParser() *CurrencyParser {
	mockData := map[string]Currency{
		"USD": {isoCode: "USD", precision: 2},
		"GBP": {isoCode: "GBP", precision: 2},
		"JPY": {isoCode: "JPY", precision: 0},
	}

	// There should never be any errors here in theory, so we can omit the second returned variable.
	parser, _ := NewCurrencyParser(&TestCurrencyData{input: mockData})

	return parser
}
