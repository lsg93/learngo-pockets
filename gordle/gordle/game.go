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
	reader     *bufio.Reader
	dictionary []string
	wordLength int
}

var (
	AskErrorInvalidInput          = errors.New("The input given is invalid")
	AskErrorNoInput               = errors.New("There needs to be input")
	AskErrorOutsideCharacterLimit = errors.New("Your guess should be X characters long")
	AskErrorNotInDictionary       = errors.New("This isn't a word")
)

var DefaultWordLength = 5

func New(options ...GameOption) *game {
	g := &game{
		input:      os.Stdin,
		wordLength: DefaultWordLength,
		dictionary: []string{"abcde", "hello", "你好你好好"},
	}

	for _, option := range options {
		option(g)
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

	if len(guess) != g.wordLength {
		return AskErrorOutsideCharacterLimit
	}

	if !slices.Contains(g.dictionary, string(guess)) {
		return AskErrorNotInDictionary
	}

	return nil
}

func (g *game) ask() (guess []rune, err error) {
	input, _, err := g.reader.ReadLine()

	str := strings.TrimSpace(string(input))

	if err != nil {
		if err == io.EOF {
			return nil, AskErrorNoInput
		}
		return nil, err
	}

	return []rune(str), nil
}

func (g *game) Play() {
	fmt.Println("Welcome to Gordle!")
	fmt.Println("Enter a guess:")
	g.ask()
}
