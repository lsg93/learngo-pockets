package money

import (
	"strings"
	"testing"
)

func TestCurrencyDataCanBeLoadedFromJSONInput(t *testing.T) {

	jsonStub := `{"GBP": {"decimals": 2}}`
	data := NewJSONCurrencyData(strings.NewReader(jsonStub))

	list, gotError := data.getCurrencies()

	if gotError != nil {
		t.Fatalf("An error %v occurred when none was expected.", gotError)
	}

	_, exists := list["GBP"]

	if !exists {
		t.Errorf("The list of currencies returned from input did not include the expected data.")
	}
}

func TestJSONCurrencyDataFailsToLoad(t *testing.T) {
	type testCase struct {
		input         string
		expectedError error
	}

	testCases := map[string]testCase{
		"With empty JSON data provided": {
			input:         `{}`,
			expectedError: errCurrencyListEmpty,
		},
		"With malformed JSON data provided": {
			input:         `{"abc" : {"def" : "ghi"}}`,
			expectedError: errCurrencyValidationError,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			data := NewJSONCurrencyData(strings.NewReader(tc.input))
			_, gotError := data.getCurrencies()

			if gotError == nil {
				t.Fatalf("No errors occurred, even though one was expected")
			}

			if gotError != tc.expectedError {
				t.Errorf("An error %v occurred, but it was not the expected error %v", gotError, tc.expectedError)
			}
		})
	}
}
