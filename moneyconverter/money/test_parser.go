package money

type TestCurrencyData struct {
	input CurrencyList
}

func (tcd *TestCurrencyData) getCurrencies() (map[string]Currency, error) {
	return tcd.input, nil
}

func setupTestCurrencyParser() *CurrencyParser {
	mockData := map[string]Currency{
		"USD": {ISOCode: "USD", Precision: 2},
		"GBP": {ISOCode: "GBP", Precision: 2},
		"JPY": {ISOCode: "JPY", Precision: 0},
	}

	// There should never be any errors here in theory, so we can omit the second returned variable.
	parser, _ := NewCurrencyParser(&TestCurrencyData{input: mockData})

	return parser
}
