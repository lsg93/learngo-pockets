package gordle

import (
	"errors"
	"strings"
)

// Default corpus taken from https://gist.github.com/dracos/dd0668f281e685bad51479e5acaadb93#file-valid-wordle-words-txt

var (
	CorpusEmptyError    = errors.New("Provided corpus is empty.")
	CorpusNotFoundError = errors.New("Corpus not found at provided path.")
)

type corpus struct {
	fs FileSystem
}

func (c *corpus) loadCorpus(path string) ([]string, error) {
	if !c.fs.FileExists(path) {
		return nil, CorpusNotFoundError
	}

	bytes, err := c.readCorpus(path)

	if err != nil {
		return nil, err
	}

	return strings.Fields(string(bytes)), err
}

func (c *corpus) readCorpus(path string) ([]byte, error) {
	bytes, err := c.fs.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if len(bytes) == 0 {
		return nil, CorpusEmptyError
	}

	return bytes, err
}
