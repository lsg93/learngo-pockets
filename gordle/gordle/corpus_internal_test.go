package gordle

import (
	"slices"
	"strings"
	"testing"
)

// A simple interface for file system operations to aid testing.

type testFs struct {
	fileExists   bool
	fileContents []string
}

func (fs *testFs) FileExists(path string) bool {
	return fs.fileExists
}

func (fs *testFs) ReadFile(path string) ([]byte, error) {
	bs := []byte(strings.Join(fs.fileContents, "\n"))

	return bs, nil
}

func TestLoadsCorpusSucessfully(t *testing.T) {
	fileContents := []string{"word1", "word2"}

	c := &corpus{fs: &testFs{fileExists: true, fileContents: fileContents}}
	gotResult, gotError := c.loadCorpus("validcorpus.txt")

	if gotError != nil {
		t.Errorf("An error %v occurred when none was expected", gotError)
	}

	if !slices.Equal(gotResult, fileContents) {
		t.Errorf("The slice returned from the loaded corpus : %+q did not match the expected result %+q", gotResult, fileContents)
	}

}

func TestLoadingCorpusFailsIfAnInvalidPathIsProvided(t *testing.T) {
	c := &corpus{fs: &testFs{fileExists: false}}
	gotResult, gotError := c.loadCorpus(".//file.txt")

	if gotResult != nil {
		t.Errorf("loading corpus returned data when it should have errored.")
	}

	if gotError != CorpusNotFoundError {
		t.Errorf("Loading corpus errored, but the error %v was not the expected error %v.", gotResult, CorpusNotFoundError)
	}
}

func TestLoadingCorpusFailsIfFileIsEmpty(t *testing.T) {
	c := &corpus{fs: &testFs{fileExists: true, fileContents: []string{}}}
	gotResult, gotError := c.loadCorpus("emptyfile.txt")

	if gotResult != nil {
		t.Errorf("Loading corpus returned data when it should have errored.")
	}

	if gotError != CorpusEmptyError {
		t.Errorf("Loading corpus errored, but the error %v was not the expected error %v.", gotResult, CorpusEmptyError)
	}
}
