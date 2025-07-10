package gordle_test

import (
	"learngo-pockets/gordle/gordle"
	"testing"
)

// Move this to an options test.
func TestGameWordLengthOption(t *testing.T) {
	type testCase struct {
		length     int
		wantResult int
	}

	testCases := map[string]testCase{
		"Length of < 2 is ignored":          {length: 1, wantResult: gordle.DefaultWordLength},
		"Given length is present on struct": {length: 6, wantResult: 6},
	}

	for _, tc := range testCases {
		g := gordle.New(
			gordle.WithWordLength(tc.length),
		)

		// TODO - this test relies on a getter not in use anywhere else, which is probably a code smell
		// When more behaviour exists, this needs to be refactored to look @ the behaviour rather than testing against unexposed internals.
		wordLength := g.WordLength()

		if wordLength != tc.wantResult {
			t.Errorf("The expected word length for this game should be %d, but it is currently %d.", tc.wantResult, wordLength)
		}
	}
}

// func TestGameAskHandlesInput(t *testing.T) {
// 	// cases
// 	// handles ascii
// 	// handles unicode
// }

// func TestGameAskHandlesErroneousInput(t *testing.T) {
// 	// errors if empty
// 	// errors if more than given length
// }
