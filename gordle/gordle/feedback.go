package gordle

/**
	Wordle operates on a two-pass algorithm.
	We first analyse the 'correct' characters.
	Then, we analyse the characters that are either absent, or in the wrong position.
	Doing it any other way has pitfalls - the logic of whether a character is wrong/absent gets hairy.
**/

// spew.Dump(map[string]interface{}{
// 	"guess char":    string(char),
// 	"solution char": string(solution[i]),
// 	"position":      i,
// })

import (
	"slices"
)

func computeFeedback(guess []rune, solution []rune) []hint {

	hints := make([]hint, len(solution))
	// We store whether a character has been processed or not as a simple boolean at its respective index in a slice.
	computed := make([]bool, len(guess))

	// Iterate over guess to find correct and absent characters.
	for i, char := range guess {
		if !slices.Contains(solution, char) {
			hints[i] = hint{character: char, status: Absent}
			computed[i] = true
		} else if char == solution[i] {
			// Create 'correct' hint at position in index
			hints[i] = hint{character: char, status: CorrectPosition}
			computed[i] = true
		}
	}

	// Iterate over guess again to find characters that exist, but are in the wrong position.
	// We can be confident that the characters in this loop exist in the solution, as we've handled absents above.

	// I'm sure this can fall over - it seems too simple.
	for i, char := range guess {
		// No need to worry about characters that have already been processed.
		if computed[i] {
			continue
		}
		hints[i] = hint{character: char, status: WrongPosition}
	}

	return hints
}
