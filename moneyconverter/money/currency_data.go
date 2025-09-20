package money

import (
	"encoding/json"
	"io"
)

type JSONCurrency struct {
	ISOCode   string
	Precision byte `json:"decimals"`
}

type CurrencyData interface {
	getCurrencies() (CurrencyList, error)
}

type JSONCurrencyData struct {
	input io.Reader
}

func NewJSONCurrencyData(input io.Reader) *JSONCurrencyData {
	return &JSONCurrencyData{input: input}
}

func (jcd *JSONCurrencyData) getCurrencies() (CurrencyList, error) {
	var decoded map[string]JSONCurrency
	err := json.NewDecoder(jcd.input).Decode(&decoded)

	if err != nil {
		return nil, err
	}

	list := make(CurrencyList)

	for code, currency := range decoded {
		list[code] = Currency{isoCode: code, precision: currency.Precision}
	}

	return list, nil
}
