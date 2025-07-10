package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type game struct {
	input      io.Reader
	reader     *bufio.Reader
	wordLength int
}

var DefaultWordLength = 5

func New(options ...GameOption) *game {
	g := &game{
		input:      os.Stdin,
		wordLength: DefaultWordLength,
	}

	for _, option := range options {
		option(g)
	}

	g.reader = bufio.NewReader(g.input)

	return g
}

func (g *game) WordLength() int {
	return g.wordLength
}

func (g *game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Enter a guess:")
}
