package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load bookworms : ")
		os.Exit(1)
	}
	fmt.Println(bookworms)
}
