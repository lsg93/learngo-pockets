package money

import "testing"

type testCase struct {
	input          []Decimal
	expectedResult Decimal
}

func TestMultipliesSuccessfully(t *testing.T) {
	testCases := map[string]testCase{
		"When 2 decimals have equal precision": {
			input:          []Decimal{{integer: 123, precision: 2}, {integer: 125, precision: 2}},
			expectedResult: Decimal{integer: 15375, precision: 4},
		},
		"When 2 decimals have similar precision": {
			input:          []Decimal{{integer: 123, precision: 2}, {integer: 15, precision: 1}},
			expectedResult: Decimal{integer: 1845, precision: 3},
		},
		"When one decimal has 0 precision, i.e. an integer.": {
			input:          []Decimal{{integer: 100, precision: 0}, {integer: 175, precision: 2}},
			expectedResult: Decimal{integer: 175, precision: 0},
		},
		"With an uninitialised argument": {
			input:          []Decimal{{integer: 100, precision: 0}, {}},
			expectedResult: Decimal{},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult := multiply(tc.input[0], tc.input[1])
			if gotResult != tc.expectedResult {
				t.Fatalf("The returned result %v was not equal to the expected result %v", gotResult, tc.expectedResult)
			}
		})
	}

}
