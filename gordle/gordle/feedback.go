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
	for i, gChar := range guess {
		// No need to worry about characters that have already been processed.
		if computed[i] {
			continue
		}

		// We want to ensure that we don't mark something as being in the wrong position just because it exists in the solution.
		// We need to be confident that we're only marking as the wrong solution if it doesnt exist anywhere further in the slice.
		// We look at the remaining characters of a solution, and check how many times our guess character occurs.
		// If it has less occurences than the remaining chars, then we can safely mark it as being in the wrong position, otherwise we can call it absent.
		// Good ref here - https://github.com/jedib0t/go-wordle/blob/main/wordle/attempt.go

		// occurences := strings.Count(string(gChar), remainingSolutionChars(solution))

	}

	return hints
}

func remainingSolutionChars(solution []rune) int {

}
