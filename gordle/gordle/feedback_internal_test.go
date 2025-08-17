package gordle

import (
	"testing"
)

// For this
// Let's test that:

// "all characters in guess are in solution, but @ wrong position",
// "all characters in guess are in solution, at right positions"
// "some characters not found in solution, some in wrong position",
// "some characters in wrong position, some in right position",
// "handles scenarios when a character is duplicated - one is absent, and one is in wrong position"
// "handles scenarios when a character is duplicated - one is wrong position, and one is in right position"
// "handles scenarios when a character is duplicated - both in right position"

func TestFeedbackServiceGeneratesHintsBasedOnGuess(t *testing.T) {
	type testCase struct {
		solution   string
		guess      string
		wantResult []hint
	}

	testCases := map[string]testCase{
		"test case 1": {
			solution:   "peers",
			guess:      "steer",
			wantResult: makeHints("steer", []hintStatus{WrongPosition, Absent, CorrectPosition, WrongPosition, WrongPosition}),
		},
		"test case 2": {
			solution:   "boost",
			guess:      "robot",
			wantResult: makeHints("robot", []hintStatus{Absent, CorrectPosition, WrongPosition, WrongPosition, CorrectPosition}),
		},
		"test case 3": {
			solution:   "apple",
			guess:      "paper",
			wantResult: makeHints("paper", []hintStatus{WrongPosition, WrongPosition, CorrectPosition, WrongPosition, Absent}),
		},
		"test case 4": {
			solution:   "apple",
			guess:      "papal",
			wantResult: makeHints("papal", []hintStatus{WrongPosition, WrongPosition, CorrectPosition, Absent, WrongPosition}),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			for i, gotHint := range gotResult {
				if compareHint(tc.wantResult[i], gotHint) == false {
					t.Errorf("Hint %v received from feedback was different to expected result %v", gotHint, tc.wantResult[i])
				}
			}
		})
	}
}

func makeHints(guess string, statuses []hintStatus) []hint {
	str := []rune(guess)
	hints := []hint{}
	for i, status := range statuses {
		hints = append(hints, hint{character: str[i], status: hintStatus(status)})
	}
	return hints
}

func compareHint(hint1 hint, hint2 hint) bool {
	if hint1.status != hint2.status || hint1.character != hint2.character {
		return false
	}
	return true
}
