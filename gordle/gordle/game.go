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
	solution   string
	dictionary []string
}

var (
	TTYGreen  = "\033[32m"
	TTYYellow = "\033[33m"
)

var (
	AskErrorInvalidInput          = errors.New("The input given is invalid")
	AskErrorNoInput               = errors.New("There needs to be input")
	AskErrorOutsideCharacterLimit = errors.New("Your guess should be X characters long")
	AskErrorNotInDictionary       = errors.New("This isn't a word")
)

func New(output io.Writer, options ...GameOption) *game {
	fmt.Println("DEBUG: New function started!")

	g := &game{
		input:      os.Stdin,
		output:     output,
		guesses:    6, // Could make this an option, but it's not wordle if it has more/less possible guesses!
		dictionary: []string{"abcde", "hello", "你好你好好"},
	}

	fmt.Println("2")

	fmt.Fprintf(os.Stdout, "DEBUG: Initial g.input type: %T\n", g.input)

	fmt.Println("3")

	for _, option := range options {
		err := option(g)
		if err != nil {
			fmt.Println(err.Error())
			output.Write([]byte(err.Error()))
		}
	}

	fmt.Fprintf(os.Stdout, "DEBUG: New g.input type: %T\n", g.input)

	g.reader = bufio.NewReader(g.input)
	fmt.Println("DEBUG: New function is about to return!")

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
	fmt.Println("making guess")
	input, _, err := g.reader.ReadLine()

	str := strings.TrimSpace(string(input))
	// fmt.Println("guess was :" + str)
	// panic("stop execution")

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
	g.output.Write([]byte("Welcome to Gordle!"))
	g.output.Write([]byte("Enter a guess:"))

	for i := 0; i < g.guesses; i++ {
		guess, err := g.ask()

		if err != nil {
			fmt.Println(err.Error())
			g.output.Write([]byte(err.Error()))
			// i-- // decrement loop
			continue
		}

		if string(guess) == g.solution {
			g.output.Write([]byte("You won!"))
		} else {
			g.output.Write([]byte(string(guess)))
		}
	}
}
