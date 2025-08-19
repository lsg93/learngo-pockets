package gordle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"unicode"
)

type game struct {
	input      io.Reader
	output     io.Writer
	reader     *bufio.Reader
	guesses    int
	solution   []rune
	dictionary []string
}

var (
	GordleOptionErrorNoSolution      = errors.New("A solution must be provided to play Gordle")
	GordleOptionErrorInvalidSolution = errors.New("The given solution must be X characters.")
)

var (
	AskErrorInvalidInput          = errors.New("The input given is invalid")
	AskErrorNoInput               = errors.New("There needs to be input")
	AskErrorOutsideCharacterLimit = errors.New("Your guess should be X characters long")
	AskErrorNotInDictionary       = errors.New("This isn't a word")
)

func New(output io.Writer, options ...GameOption) *game {

	g := &game{
		input:      os.Stdin,
		output:     output,
		guesses:    6, // Could make this an option, but it's not wordle if it has more/less possible guesses!
		dictionary: []string{"abcde"},
	}

	for _, option := range options {
		err := option(g)
		if err != nil {
			fmt.Println(err.Error())
			output.Write([]byte(err.Error() + "\n"))
			os.Exit(1)
		}
	}

	if len(g.solution) == 0 {
		output.Write([]byte(GordleOptionErrorNoSolution.Error() + "\n"))
		os.Exit(1)
	}

	g.reader = bufio.NewReader(g.input)

	return g
}

func isPurelyAlphabetical(guess []rune) bool {
	for _, r := range guess {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (g *game) validateGuess(guess []rune) error {

	if !isPurelyAlphabetical(guess) {
		return AskErrorInvalidInput
	}

	if len(guess) != len(g.solution) {
		return AskErrorOutsideCharacterLimit
	}

	if !slices.Contains(g.dictionary, string(guess)) {
		return AskErrorNotInDictionary
	}

	return nil
}

func (g *game) ask() (guess []rune, err error) {
	g.output.Write([]byte("Enter a guess:" + "\n"))
	input, _, err := g.reader.ReadLine()

	str := strings.TrimSpace(string(input))

	if err != nil {
		if err == io.EOF {
			return nil, AskErrorNoInput
		}
		return nil, err
	}

	vErr := g.validateGuess([]rune(str))

	if vErr != nil {
		return nil, vErr
	}

	return []rune(str), nil
}

func (g *game) Play() {
	g.output.Write([]byte("Welcome to Gordle!" + "\n"))

	for i := 0; i < g.guesses; i++ {
		guess, err := g.ask()

		if err != nil {
			g.output.Write([]byte(err.Error() + "\n"))
			i-- // decrement loop
			continue
		}

		feedback := computeFeedback(guess, g.solution)

		if string(guess) == string(g.solution) {
			g.output.Write([]byte("You won!"))
			return
		} else {
			g.output.Write([]byte(feedback.String() + "\n"))
		}

	}
}
