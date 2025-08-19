package gordle

import (
	"io"
)

type GameOption func(*game) error

func WithDictionary(dictionary []string) GameOption {
	return func(g *game) error {
		g.dictionary = dictionary
		return nil
	}
}

func WithInput(input io.Reader) GameOption {
	return func(g *game) error {
		g.input = input
		return nil
	}
}

func WithSolution(solution string) GameOption {
	return func(g *game) error {
		if len(solution) < 2 {
			return GordleOptionErrorInvalidSolution
		}
		g.solution = []rune(solution)
		return nil
	}
}
