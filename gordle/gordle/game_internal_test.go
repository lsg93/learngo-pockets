package gordle

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestGameAskInput(t *testing.T) {
	type testCase struct {
		wantError  error
		wantResult string
		input      string
	}

	testCases := map[string]testCase{
		"handles ascii text":            {wantError: nil, wantResult: "abcde", input: "abcde"},
		"handles unicode characters":    {wantError: nil, wantResult: "你好", input: "你好"},
		"trims input on left side":      {wantError: nil, wantResult: "abcd", input: " abcd"},
		"trims input on right side":     {wantError: nil, wantResult: "abcd", input: "abcd "},
		"trims and handles empty input": {wantError: AskErrorNoInput, wantResult: "", input: " "},
		"handles empty input":           {wantError: AskErrorNoInput, wantResult: "", input: ""},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			reader := bytes.NewBufferString(tc.input)
			g := New(os.Stdout, WithInput(reader))
			guess, gotError := g.ask()

			if string(guess) != tc.wantResult {
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
			g := New(os.Stdout)
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
