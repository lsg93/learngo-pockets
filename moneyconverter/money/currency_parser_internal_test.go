package money

import (
	"testing"
)

// Takes in input and returns a struct with expected data accessible somehow.
// Fails with malformed data?

type testCurrencyService struct {
	currencies map[string]Currency
}

func (cs *testCurrencyService) getCurrencies() {
	return cs.currencies
}

func TestCurrencyParserInitialisesSuccessfullyWithExpectedData(t *testing.T) {

	stub := map[string]Currency{
		"GBP": Currency{isoCode: "GBP", decimal: 2},
	}

	service := &testCurrencyService{currencies: stub}

	parser, gotError := NewCurrencyParser(service)

	if gotError != nil {
		t.Errorf("An error %v occurred when none was expected", gotError)
	}

	_, exists := parser.currencies["GBP"]

	if !exists {
		t.Errorf("The currency parser initialised, but the expected data was not present.")
	}
}

// func TestCurrencyParserFailsToInitialise() {
// 	type testCase struct {
// 		input         string
// 		expectedError error
// 	}

// 	testCases := map[string]testCase{
// 		"With no data provided" : {
// 			input : `{}`,
// 			expectedError: ,
// 		}
// 		"With malformed data provided" : {
// 			input : `{}`,
// 			expectedError: ,
// 		}

// 	}
// }
