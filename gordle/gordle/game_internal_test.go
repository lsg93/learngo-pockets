package gordle

import (
	"bytes"
	"errors"
	"os"
	"slices"
	"testing"
)

func TestGameSetupLoadsDefaultCorpusIfNoneProvidedAsOption(t *testing.T) {
	defaultCorpus := []string{"salet", "audio", "light"}
	fs := &testFs{fileExists: true, fileContents: defaultCorpus}
	g := New(os.Stdout, WithFileSystem(fs))

	if !slices.Equal(defaultCorpus, g.dictionary) {
		t.Errorf("Since no dictionary was provided, the dictionary used for the game should be the provided dictionary.")
	}
}

func TestGameSetupPicksRandomSolutionFromDictionaryIfNoneProvidedAsOption(t *testing.T) {
	dictionary := []string{"salet", "audio", "light"}
	g := New(os.Stdout, WithDictionary(dictionary))

	if !slices.Contains(dictionary, string(g.solution)) {
		t.Errorf("Since no solution was provided, the solution should have been selected from the provided dictionary.")
	}
}

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
			g := New(&bytes.Buffer{}, WithDictionary([]string{"tests"}), WithInput(reader))
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

func TestGameAskValidatesInputProperly(t *testing.T) {
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
			g := New(os.Stdout, WithDictionary([]string{"hello", "你好你好好"}))
			gotError := g.validateGuess([]rune(tc.guess))

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
