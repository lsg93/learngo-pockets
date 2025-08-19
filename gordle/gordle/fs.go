package gordle

import "os"

// A simple interface for file system operations to aid testing.

type FileSystem interface {
	FileExists(path string) bool
}

type gordleFs struct{}

func (fs *gordleFs) FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

type testFs struct {
	fileExists bool
}

func (fs *testFs) FileExists(path string) bool {
	return fs.fileExists
}
