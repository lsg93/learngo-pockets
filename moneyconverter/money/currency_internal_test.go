package money

import "testing"

func TestParseCurrency(t *testing.T) {

	type testCase struct {
		input          string
		expectedResult Currency
		expectedError  error
	}

	testCases := map[string]testCase{
		"Throws error when non-existent ISO 4217 currency code is supplied": {
			input:          "GBK",
			expectedResult: Currency{},
			expectedError:  CurrencyParseInvalidCodeError,
		},
		"Throws error when currency code is not English": {
			input:          "💰💰💰",
			expectedResult: Currency{},
			expectedError:  CurrencyParseInvalidCodeError,
		},
		"Supports lowercase": {
			input:          "gbp",
			expectedResult: Currency{isoCode: "GBP", precision: 2},
			expectedError:  nil,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult, gotError := ParseCurrency(tc.input)

			if tc.expectedError != gotError {
				t.Errorf("The error that occurred was not the expected error")
			}

			if tc.expectedResult != gotResult {
				t.Errorf("The expected result was not the result received")
			}
		})
	}
}
