package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"testing"
)

func TestConvert(t *testing.T) {

	type testCase struct {
		convertFrom    money.Amount
		convertTo      string
		expectedError  error
		expectedAmount money.Amount
	}

	testCases := map[string]testCase{
		"1 GBP to USD value": {
			convertFrom:    money.NewAmount(1, "GBP"),
			convertTo:      "USD",
			expectedError:  nil,
			expectedAmount: money.NewAmount(1.33, "USD"),
		},
		"19.99 USD to GBP value": {
			convertFrom:    money.NewAmount(19.99, "USD"),
			convertTo:      "USD",
			expectedError:  nil,
			expectedAmount: money.NewAmount(14.80, "GBP"),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotAmount, gotErr := money.Convert(tc.convertFrom, money.ParseCurrency(tc.convertTo))

			if tc.expectedError == nil && gotErr != nil {
				t.Errorf("An error %v was returned and none was expected.", gotErr)
			}

			if tc.expectedError != nil && gotErr != tc.expectedError {
				t.Errorf("An error %v was returned but error %v was expected.", gotErr, tc.expectedError)
			}

			if !gotAmount.Equals(tc.expectedAmount) {
				t.Errorf("Conversion failed - expected output was %d, %d given", tc.expectedAmount, gotAmount)
			}
		})
	}
}
