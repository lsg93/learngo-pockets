package money

import (
	"strings"
	"testing"
)

// Takes in input and returns a struct with expected data accessible somehow.
// Fails with malformed data?

func TestCurrencyParserInitialisesSuccessfullyWithJSONData(t *testing.T) {

	jsonStub := `{"GBP": {"decimals": 2}}`
	data := NewJSONCurrencyData(strings.NewReader(jsonStub))

	parser, gotError := NewCurrencyParser(data)

	if gotError != "" {
		t.Fatalf("An error %v occurred when none was expected", gotError)
	}

	_, exists := parser.currencies["GBP"]

	if !exists {
		t.Errorf("The currency parser initialised, but the expected data was not present.")
	}
}

func TestCurrencyParserFailsToInitialise(t *testing.T) {
	type testCase struct {
		input         string
		expectedError error
	}

	testCases := map[string]testCase{
		"With empty JSON data provided": {
			input:         `{}`,
			expectedError: CurrencyDataNotLoadedError,
		},
		"With malformed JSON data provided": {
			input:         `{"abc" : {"def" : "ghi"}}`,
			expectedError: CurrencyDataReadError,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			data := NewJSONCurrencyData(strings.NewReader(tc.input))
			parser, gotError := NewCurrencyParser(data)

			if gotError != "" {
				return
			}

		})
	}
}
