package gordle

import (
	"testing"
)

func TestLoadingCorpusFailsIfAnInvalidPathIsProvided(t *testing.T) {
	// throw CorpusNotFoundError
	c := &corpus{}
	gotResult, gotError := c.loadCorpus(&testFs{fileExists: false})

	if gotResult != nil {
		t.Errorf("loading corpus returned data when it should have errored.")
	}

	if gotError != CorpusNotFoundError {
		t.Errorf("loading corpus errored, but the error was not the expected error.")
	}
}

func TestLoadingCorpusFailsIfFileIsEmpty(t *testing.T) {
	// throw CorpusEmptyError
	t.Skip("TODO")
}

func TestLoadsCorpusSucessfully(t *testing.T) {
	// probably have to mock io reader somehow
	// check that 'file' property points to a byte slice?
	t.Skip("TODO")
}

func TestConvertsLoadedCorpusIntoSliceOfStrings(t *testing.T) {
	// probably have to mock io reader again?
	// check that 'file' property when converted to string = returned slice?
	t.Skip("TODO")
}
