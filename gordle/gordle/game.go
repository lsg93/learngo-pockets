package gordle

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"slices"
	"strings"
	"unicode"
)

type game struct {
	fs         FileSystem
	input      io.Reader
	output     io.Writer
	reader     *bufio.Reader
	guesses    int
	solution   []rune
	dictionary []string
}

const DEFAULT_CORPUS_PATH = "./gordle/corpus.txt"

var (
	GordleOptionNoDictionary         = errors.New("A dictionary must be provided to play Gordle.")
	GordleOptionErrorNoSolution      = errors.New("A solution must be provided to play Gordle.")
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
		input:   os.Stdin,
		output:  output,
		fs:      &gordleFs{},
		guesses: 6, // Could make this an option, but it's not wordle if it has more/less possible guesses!
	}

	for _, option := range options {
		err := option(g)
		if err != nil {
			output.Write([]byte(err.Error() + "\n"))
			os.Exit(1)
		}
	}

	if len(g.dictionary) == 0 {
		corpusLoader := &corpus{fs: g.fs}
		dict, err := corpusLoader.loadCorpus(DEFAULT_CORPUS_PATH)

		if err != nil {
			output.Write([]byte(err.Error() + "\n"))
			os.Exit(1)
		}

		g.dictionary = dict
	}

	if len(g.solution) == 0 {
		if len(g.dictionary) == 0 {
			output.Write([]byte(GordleOptionNoDictionary.Error() + "\n"))
			os.Exit(1)
		}
		index := rand.Intn(len(g.dictionary))
		g.solution = []rune(g.dictionary[index])
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

	return []rune(str), nil
}

func (g *game) Play() {
	g.output.Write([]byte("Welcome to Gordle!" + "\n"))

	for i := 0; i < g.guesses; i++ {
		guess, err := g.ask()

		vErr := g.validateGuess(guess)

		if vErr != nil {
			g.output.Write([]byte(err.Error() + "\n"))
			continue
		}

		if err != nil {
			g.output.Write([]byte(err.Error() + "\n"))
			i-- // decrement loop
			continue
		}

		feedback := computeFeedback(guess, g.solution)

		g.output.Write([]byte(feedback.String() + "\n"))

		if string(guess) == string(g.solution) {
			g.output.Write([]byte("You won!"))
			return
		}
	}
}
