package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"testing"
)

func TestConvertsSuccesfully(t *testing.T) {

	type testCase struct {
		convertAmount  string
		convertFrom    string
		convertTo      string
		rate           string
		expectedResult money.Amount
	}

	testCases := map[string]testCase{
		"1 GBP to USD value": {
			convertAmount:  "1",
			convertFrom:    "GBP",
			convertTo:      "USD",
			rate:           "1.3455", // Rate as of 25/09/25
			expectedResult: money.Amount{},
		},
		"19.99 USD to GBP value": {
			convertAmount:  "19.99",
			convertFrom:    "USD",
			convertTo:      "GBP",
			expectedResult: money.Amount{},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			convertFrom, convertTo := money.SetupTestConversionData(t, tc.convertAmount, tc.convertFrom, tc.convertTo)
			converter := money.SetupTestConverter(t, tc.rate)

			gotResult, gotError := converter.Convert(convertFrom, convertTo)

			if gotError != nil {
				t.Fatalf("An error %v was returned and none was expected.", gotError)
			}

			if gotResult != tc.expectedResult {
				t.Errorf("The amount received %v and the expected amount %v were not equal", gotResult, tc.expectedResult)
			}
		})
	}
}
