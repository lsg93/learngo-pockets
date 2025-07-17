package gordle

import (
	"errors"
	"io"
)

var (
	GordleOptionErrorInvalidSolution = errors.New("The given solution must be X characters.")
)

type GameOption func(*game) error

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
		g.solution = solution
		return nil
	}
}
