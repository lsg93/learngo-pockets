package main

import (
	"flag"
	"fmt"
	"os"
)

func displayBooks(books []Book) {
	if len(books) == 0 {
		fmt.Println("No common books were found.")
		return
	}

	fmt.Println("Here are the common books:")
	for _, book := range books {
		fmt.Printf("- %s by %s\n", book.Name, book.Author)
	}
}

func main() {
	var path string
	flag.StringVar(&path, "bookwormPath", "testdata/bookworms.json", "The relative path to bookworm data stored in JSON format.")
	flag.Parse()

	bookworms, err := loadBookworms(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load bookworms : ")
		os.Exit(1)
	}
	cb := findCommonBooks(bookworms)
	displayBooks(cb)
}
