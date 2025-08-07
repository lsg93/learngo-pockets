package gordle

import (
	"slices"

	"github.com/davecgh/go-spew/spew"
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
	seen := []int{}

	for i, char := range guess {
		if char == solution[i] {
			// Create 'correct' hint at position in index
			hints[i] = hint{character: char, status: CorrectPosition}
			// Mark character as "seen" so it is ignored/unused in next pass.
			seen = append(seen, i)
		}
	}

	for i, char := range guess {
		spew.Dump("second pass - character is " + string(char))
		if slices.Contains(seen, i) {
			spew.Dump("ignore character " + string(char) + " it has already been seen.")
			continue
		}

		if slices.Contains(solution, char) {
			// We know that the character in question is in the wrong position here.
			hints[i] = hint{character: char, status: WrongPosition}
			// If there are more occurences of this character in the solution, then we continue on without setting it to "seen".
			if charOccursInSolutionAgain(char, solution) {
				continue
			}
		} else {
			hints[i] = hint{character: char, status: Absent}
		}

		seen = append(seen, i)
	}

	return hints
}

func charOccursInSolutionAgain(char rune, solution []rune) bool {

	count := 0
	for _, r := range solution {
		if r == char {
			count++
		}
	}

	return count > 1
}
