package gordle

import (
	"strings"
)

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

func computeFeedback(guess []rune, solution []rune) []hint {

	hints := make([]hint, len(solution))
	used := make([]bool, len(solution))

	// Iterate over guess to find correct characters.
	for i, char := range guess {
		if solution[i] == char {
			hints[i] = hint{character: char, status: CorrectPosition}
		}
	}

	// Iterate over guess again to find characters that exist, but are in the wrong position.
	for i, char := range guess {
		if hints[i].status == CorrectPosition {
			continue
		}
		for j, solutionChar := range solution {
			if !used[j] && char == solutionChar {
				hints[i] = hint{character: char, status: WrongPosition}
				used[j] = true
				break
			}
		}
	}

	for i := range hints {
		if hints[i].status == Undefined {
			hints[i] = hint{character: guess[i], status: Absent}
		}
	}

	return hints
}

func getCharOccurences(charToCount rune, slice []rune) int {
	return strings.Count(string(slice), string(charToCount))
}
