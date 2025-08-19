package gordle

import (
	"errors"
	"io"
)

// Default corpus taken from https://gist.github.com/dracos/dd0668f281e685bad51479e5acaadb93#file-valid-wordle-words-txt

var (
	CorpusNotFoundError = errors.New("Corpus not found at provided path.")
)

type corpus struct {
	path string
}

func (c *corpus) loadCorpus(fs FileSystem) ([]string, error) {
	if !fs.FileExists(c.path) {
		return nil, CorpusNotFoundError
	}

	return []string{}, nil
}

// func (c *corpus)

func (c *corpus) loadFile(fs FileSystem, reader io.Reader) {
	// reader.Read()
}
