package money

import (
	"testing"
)

func TestParseCurrencySuccesfully(t *testing.T) {

	type testCase struct {
		input          string
		expectedResult Currency
	}

	testCases := map[string]testCase{
		"Supports typical input": {
			input:          "USD",
			expectedResult: Currency{ISOCode: "USD", Precision: 2},
		},
		"Supports lowercase": {
			input:          "gbp",
			expectedResult: Currency{ISOCode: "GBP", Precision: 2},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			parser := setupTestCurrencyParser()
			gotResult, gotError := parser.ParseCurrency(tc.input)

			if gotError != nil {
				t.Errorf("An error %s occurred when none was expected.", gotError)
			}

			if tc.expectedResult != gotResult {
				t.Errorf("The expected result %v was not the result received : %v", tc.expectedResult, gotResult)
			}
		})
	}
}
func TestParseCurrencyInputValidation(t *testing.T) {

	type testCase struct {
		input         string
		expectedError error
	}

	testCases := map[string]testCase{
		"Throws error when currency code is not English": {
			input:         "💰💰💰",
			expectedError: errCurrencyInputValidation,
		},
		"Throws error when currency code != 3 characters": {
			input:         "ab",
			expectedError: errCurrencyInputValidation,
		},
		"Throws error with numeric input": {
			input:         "123",
			expectedError: errCurrencyInputValidation,
		},
		"Throws error when valid, but non-existent ISO 4217 currency code is supplied": {
			input:         "GBK",
			expectedError: errCurrencyNotFound,
		},
	}

	for desc, tc := range testCases {

		t.Run(desc, func(t *testing.T) {
			parser := setupTestCurrencyParser()
			_, gotError := parser.ParseCurrency(tc.input)

			if gotError == nil {
				t.Errorf("No error occurred, but the error %v was expected", tc.expectedError)
			}

			if tc.expectedError != gotError {
				t.Errorf("The expected error %v was not the error received : %v", tc.expectedError, gotError)
			}
		})
	}
}
