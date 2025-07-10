package gordle

import "io"

type GameOption func(*game)

func WithInput(input io.Reader) GameOption {
	return func(g *game) {
		g.input = input
	}
}

func WithWordLength(length int) GameOption {
	return func(g *game) {
		// Cannot be a word if length is 0 or 1
		if length < 2 {
			return
		}
		g.wordLength = length
	}
}
