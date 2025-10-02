package money

import "testing"

func TestDecimalParse(t *testing.T) {
	type testCase struct {
		expectedResult Decimal
		expectedErr    error
		input          string
	}

	testCases := map[string]testCase{
		"Handles X input": {
			input:          "123",
			expectedResult: Decimal{integer: 123, precision: 0},
			expectedErr:    nil,
		},
		"Handles X.XX input": {
			input:          "1.23",
			expectedResult: Decimal{integer: 123, precision: 2},
			expectedErr:    nil,
		},
		"Handles XX.XX input": {
			input:          "12.34",
			expectedResult: Decimal{integer: 1234, precision: 2},
			expectedErr:    nil,
		},
		"Handles 0.XX input": {
			input:          "0.12",
			expectedResult: Decimal{integer: 12, precision: 2},
			expectedErr:    nil,
		},
		"Handles X.00 input": {
			input:          "1.00",
			expectedResult: Decimal{integer: 1, precision: 0},
			expectedErr:    nil,
		},
		"Handles leading 0's": {
			input:          "01.23",
			expectedResult: Decimal{integer: 123, precision: 2},
			expectedErr:    nil,
		},
		"Handles trailing 0's": {
			input:          "32.10",
			expectedResult: Decimal{integer: 321, precision: 1},
			expectedErr:    nil,
		},
		// "Handles separators in integer parts": {
		// 	input:          "123,456.78",
		// 	expectedResult: "123.45",
		// 	expectedErr:    nil,
		// },
		// "Handles integer parts with length > 2": {
		// 	input:          "123.45",
		// 	expectedResult: "123.45",
		// 	expectedErr:    nil,
		// },
		// "Handles mantissas with length > 2": {
		// 	input:          "12.345",
		// 	expectedResult: "12.35",
		// 	expectedErr:    nil,
		// },
		// "Handles leading whitespace": {
		// 	input:          " 1.23",
		// 	expectedResult: "1.23",
		// 	expectedErr:    nil,
		// },
		// "Handles trailing whitespace": {
		// 	input:          "1.23 ",
		// 	expectedResult: "1.23",
		// 	expectedErr:    nil,
		// },
		// "Errors on empty input": {
		// 	input:          "",
		// 	expectedResult: "",
		// 	expectedErr:    DecimalParseErrorNoInput,
		// },
		// "Errors on invalid input": {
		// 	input:          "ab.cd",
		// 	expectedResult: "",
		// 	expectedErr:    DecimalParseErrorInvalidInput,
		// },
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult, gotErr := ParseDecimal(tc.input)
			if gotErr != nil && gotErr != tc.expectedErr {
				t.Errorf("An unexpected error was thrown %v", gotErr)
			}

			if gotResult != tc.expectedResult {
				t.Errorf("The returned result %v was not equal to the expected result %v", gotResult, tc.expectedResult)
			}
		})
	}
}

func TestDecimalTargetPrecisionCanBeSetSuccesfully(t *testing.T) {
	type testCase struct {
		decimal         Decimal
		targetPrecision int
		expectedResult  Decimal
	}

	testCases := map[string]testCase{
		"When precision of quantity is the same as precision of currency": {
			decimal:         Decimal{integer: 133, precision: 2},
			targetPrecision: 2,
			expectedResult:  Decimal{integer: 133, precision: 2},
		},
		"When precision of quantity is greater than precision of currency": {
			decimal:         Decimal{integer: 133, precision: 2},
			targetPrecision: 0,
			expectedResult:  Decimal{integer: 1, precision: 0},
		},
		"When precision of quantity is less than precision of currency": {
			decimal:         Decimal{integer: 133, precision: 2},
			targetPrecision: 3,
			expectedResult:  Decimal{integer: 1330, precision: 3},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.decimal.setTargetPrecision(tc.targetPrecision)
			if tc.decimal != tc.expectedResult {
				t.Errorf("The expected result %d was not equal to the returned result %d", tc.decimal, tc.expectedResult)
			}
		})
	}

}
