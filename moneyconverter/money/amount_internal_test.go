package money

import "testing"

func TestAmountInitialisesSuccessfully(t *testing.T) {

	type testCase struct {
		quantityInput  Decimal
		currencyInput  Currency
		expectedResult Amount
	}

	testCases := map[string]testCase{
		"When quantity and currency precision are the same": {
			quantityInput:  Decimal{},
			currencyInput:  Currency{},
			expectedResult: Amount{},
		},
		"When quantity and currency precision are not the same, but not a mismatch": {
			quantityInput:  Decimal{},
			currencyInput:  Currency{},
			expectedResult: Amount{},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult, gotError := NewAmount(tc.quantityInput, tc.currencyInput)

			if gotError != nil {
				t.Fatalf("An error %v was returned when none was expected", gotError)
			}

			if gotResult != tc.expectedResult {
				t.Errorf("The result returned %v was not equal to the expected result %v", gotResult, tc.expectedResult)
			}
		})
	}
}

func TestAmountFailsToInitialiseWithMismatchedPrecision(t *testing.T) {
	q := Decimal{integer: 123456, precision: 3}
	curr := Currency{ISOCode: "GBP", Precision: 2}

	_, gotError := NewAmount(q, curr)

	if gotError == nil {
		t.Fatalf("No error was returned when error %v was expected", errAmountPrecisionMismatch)
	}

	if gotError != errAmountPrecisionMismatch {
		t.Errorf("The error returned %v was not equal to the expected error %v", gotError, errAmountPrecisionMismatch)
	}
}
