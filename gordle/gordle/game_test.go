package gordle_test

import (
	"bytes"
	"fmt"
	"learngo-pockets/gordle/gordle"
	"testing"
)

type testCase struct {
	solution string
	guesses  []testValidGuess
}

type testValidGuess struct {
	guess   string
	colours []string
}

//  testCases := map[string]testCase {
// 	"user completes game after a succesful guess of the solution" : {}
// 	"user loses game after exhausting max number of guesses" : {}
// 	"user is warned about empty input"
// 	"user is warned about input that is not in the dictionary"
// 	"user is warned about numeric input"
// 	"user is warned about usage of special characters"
// }

func TestGordleGameWin(t *testing.T) {
	type testCase struct {
		solution string
		guesses  []testValidGuess
	}

	testCases := map[string]testCase{
		"user completes game after a succesful guess of the solution": {
			solution: "audio",
			guesses: []testValidGuess{
				{guess: "guess\n", colours: []string{"", gordle.TTYGreen, "", "", ""}},
				{guess: "quart\n", colours: []string{"", gordle.TTYGreen, gordle.TTYYellow, "", ""}},
				{guess: "dummy\n", colours: []string{gordle.TTYYellow, gordle.TTYGreen, "", "", ""}},
				{guess: "audit\n", colours: []string{gordle.TTYGreen, gordle.TTYGreen, gordle.TTYGreen, gordle.TTYGreen, ""}},
				{guess: "audio\n", colours: []string{gordle.TTYGreen, gordle.TTYGreen, gordle.TTYGreen, gordle.TTYGreen, gordle.TTYGreen}},
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			reader := &bytes.Buffer{}
			output := &bytes.Buffer{}

			for _, guess := range tc.guesses {
				reader.WriteString(guess.guess)
			}

			g := gordle.New(
				output,
				gordle.WithInput(reader),
				gordle.WithSolution(tc.solution),
				gordle.WithDictionary([]string{"audio", "guess", "quart", "dummy", "audit"}),
			)

			g.Play()

			fmt.Println(output)

			// Simulate terminal input somehow.
			// For each input guess, assert that the output is also valid?

		})
	}
}

func parseValidGuessOutput(output string, guess testValidGuess) string {
	// TODO

	return ""
}

// func TestGordleGameLoss(t *testing.T) {
// 	// Test Loss scenarios here.
// }

// func TestGordleGameInput() {
// 	// Test invalid feedback attempts here.
// }
