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
	GetCurrencies() (CurrencyList, error)
}

type JSONCurrencyData struct {
	input io.Reader
}

func NewJSONCurrencyData(input io.Reader) *JSONCurrencyData {
	return &JSONCurrencyData{input: input}
}

func (jcd *JSONCurrencyData) GetCurrencies() (CurrencyList, error) {
	var decoded map[string]JSONCurrency
	err := json.NewDecoder(jcd.input).Decode(&decoded)

	if err != nil {
		return nil, err
	}

	list := make(CurrencyList)

	if len(decoded) == 0 {
		return nil, errCurrencyListEmpty
	}

	for code, currency := range decoded {
		c := Currency{ISOCode: code, Precision: currency.Precision}

		err := c.validate()

		if err != nil {
			return nil, err
		}

		list[code] = c
	}

	return list, nil
}
