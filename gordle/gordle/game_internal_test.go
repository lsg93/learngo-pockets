package gordle

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// Move this to an options test.
func TestGameWordLengthOption(t *testing.T) {
	type testCase struct {
		length     int
		wantResult int
	}

	testCases := map[string]testCase{
		"Length of < 2 is ignored":          {length: 1, wantResult: DefaultWordLength},
		"Given length is present on struct": {length: 6, wantResult: 6},
	}

	for _, tc := range testCases {
		g := New(
			WithWordLength(tc.length),
		)

		wordLength := g.wordLength

		if wordLength != tc.wantResult {
			t.Errorf("The expected word length for this game should be %d, but it is currently %d.", tc.wantResult, wordLength)
		}
	}
}

func TestGameAskInput(t *testing.T) {
	type testCase struct {
		wantError error
		input     string
	}

	testCases := map[string]testCase{
		"handles ascii text":         {wantError: nil, input: "abcde"},
		"handles unicode characters": {wantError: nil, input: "你好"},
		"handles empty input":        {wantError: AskErrorNoInput, input: ""},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			reader := bytes.NewBufferString(tc.input)
			g := New(WithInput(reader))
			guess, gotError := g.ask()

			if string(guess) != tc.input {
				t.Errorf("The input string expected is not the same as the input string received.")
			}

			if tc.wantError == nil && gotError != nil {
				t.Errorf("An error occurred, but none was supposed to.")
			}

			if tc.wantError != nil && gotError != nil {
				if !errors.Is(gotError, tc.wantError) {
					t.Errorf("An error %s was thrown, but it is not the expected error %s", gotError.Error(), tc.wantError.Error())
				}
			}

		})
	}
}

func TestGameAskValdiatesInputProperly(t *testing.T) {
	type testCase struct {
		wantError error
		guess     string
	}

	testCases := map[string]testCase{
		"allows for valid ascii guesses":                             {wantError: nil, guess: "hello"},
		"allows for valid unicode guesses":                           {wantError: nil, guess: "你好你好好"},
		"rejects numbers":                                            {wantError: AskErrorInvalidInput, guess: "12345"},
		"rejects special characters":                                 {wantError: AskErrorInvalidInput, guess: "<?/->"},
		"rejects ascii input that is not equal to character limit":   {wantError: AskErrorOutsideCharacterLimit, guess: "abc"},
		"rejects unicode input that is not equal to character limit": {wantError: AskErrorOutsideCharacterLimit, guess: "你好你好你好"},
		"rejects input that is not in the dictionary":                {wantError: AskErrorNotInDictionary, guess: "notin"},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			g := New()
			gotError := g.validateGuess([]rune(tc.guess))

			if tc.wantError == nil && gotError != nil {
				fmt.Println(tc.guess)
				fmt.Println(gotError.Error())
				t.Errorf("An error occurred, but none was supposed to.")
			}

			if tc.wantError != nil && gotError != nil {
				if !errors.Is(gotError, tc.wantError) {
					t.Errorf("An error %s was thrown, but it is not the expected error %s", gotError.Error(), tc.wantError.Error())
				}
			}

		})
	}
}
