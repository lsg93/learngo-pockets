package main

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
)

type Book struct {
	Author string `json:"author"`
	Name   string `json:"name"`
}

type BookCount map[Book]uint
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bufReader := bufio.NewReaderSize(f, 1024*1024)
	var bookworms []Bookworm
	err = json.NewDecoder(bufReader).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// Find books that belong to more than one person in our list of bookworms.
func findCommonBooks(bookworms []Bookworm) []Book {
	var books []Book
	counts := getBookCounts(bookworms)

	for book, count := range counts {
		if count > 1 {
			books = append(books, book)
		}
	}

	return sortBooks(books)
}

func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Name < books[j].Name
	})
	return books
}

func getBookCounts(bookworms []Bookworm) BookCount {
	cb := make(map[Book]uint)
	for _, p := range bookworms {
		for _, book := range p.Books {
			cb[book]++
		}
	}
	return cb
}
