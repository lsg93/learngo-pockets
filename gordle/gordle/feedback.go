package gordle

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type feedback struct {
	guess    []rune
	seen     []rune
	solution []rune
	result   []hint
}

type analysedCharacter struct {
	indexes     []int
	atPositions []int
}

func computeFeedback(guess []rune, solution string) []hint {

	// using string type for dumping purposes - come back and fix this.
	feedback := make([]hint, 5)
	analysedCharacters := make(map[string]analysedCharacter)

	for i, ch := range guess {
		str := string(ch)
		v, ok := analysedCharacters[str]
		if !ok {
			analysedCharacters[str] = analysedCharacter{
				indexes:     []int{i},
				atPositions: getPositionsOfCharacterInSolution(str, solution),
			}
		} else {
			// This needs to be fixed too, its not appending properly.
			v.indexes = append(v.indexes, i)
		}
	}

	// spew.Dump("solution")
	// spew.Dump(solution)
	// spew.Dump("map")
	spew.Dump(analysedCharacters)

	// Nearly there, needs work
	for ch, analysis := range analysedCharacters {
		for _, position := range analysis.indexes {
			runeCh := []rune(ch)[0]
			feedback[position] = hint{character: runeCh, status: Absent}
			if len(analysis.atPositions) > 0 {
				for _, charIndex := range analysis.atPositions {
					if strings.Index(solution, ch) == charIndex {
						feedback[position] = hint{character: runeCh, status: CorrectPosition}
					} else {
						feedback[position] = hint{character: runeCh, status: WrongPosition}
					}
				}
			}
		}
	}

	spew.Dump(feedback)

	return []hint{}
}

func getPositionsOfCharacterInSolution(char string, solution string) []int {
	positions := []int{}
	for i, ch := range solution {
		// spew.Printf("letter %s \n", string(ch))
		if string(ch) == char {
			// spew.Printf("letter %s at index %d", string(ch), i)
			positions = append(positions, i)
		}
	}
	return positions
}

// func computeFeedback(guess []rune, solution string) []hint {
// 	feedback := &feedback{
// 		guess:    guess,
// 		seen:     []rune{},
// 		solution: []rune(solution),
// 		result:   make([]hint, 5),
// 	}

// 	spew.Dump("Solution is %s", string(feedback.solution))
// 	spew.Dump("Guess is %s", string(feedback.guess))

// 	feedback.computeAbsent()
// 	feedback.computePositions()

// 	spew.Dump("Feedback")
// 	spew.Dump(map[string]interface{}{
// 		"seen":   r2str(feedback.seen),
// 		"result": feedback.result,
// 	})

// 	return feedback.makeResult()

// }

// func (f *feedback) computeAbsent() {

// 	spew.Dump("Computing absent characters")

// 	for i, ch := range f.guess {
// 		if !slices.Contains(f.solution, ch) {
// 			f.result[i] = hint{character: ch, status: Absent}
// 			f.seen = append(f.seen, ch)
// 		}
// 	}
// }

// func (f *feedback) buildPositions() {
// 	spew.Dump("Computing character positions")

// }

// func (f *feedback) makeResult() []hint {
// 	return make([]hint, 5)
// }

func r2str(s []rune) []string {
	var ss []string
	for _, r := range s {
		ss = append(ss, string(r))
	}
	return ss
}
