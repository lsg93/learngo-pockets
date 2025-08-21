package gordle

import (
	"os"
)

type FileSystem interface {
	FileExists(path string) bool
	ReadFile(path string) ([]byte, error)
}

type gordleFs struct{}

func (fs *gordleFs) FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (fs *gordleFs) ReadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
