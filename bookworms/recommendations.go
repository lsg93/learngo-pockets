package main

import (
	"math"
	"slices"
)

// The descriptions given for certain implementation details in the book are very terse/nonexistent
// So this might not be a perfect analogue to the function shown in the book @ Listing 3.22.

func bookIntersection(bwSlice []Book, targetSlice []Book) ([]Book, []Book) {
	set := map[Book]bool

	for _, book := range bwSlice {
		set[book] = true
	}

	intersection := []Book{}
	unreadBooks := []Book{}

	for _, book := range targetSlice {
		if _, ok := set[book]; ok {
			intersection = append(intersection, book)
		} else {
			unreadBooks = append(unreadBooks, book)
		}
	}

	return intersection, unreadBooks
}

func recommendBooks(bookworm string, bookworms []Bookworm, recommendationCount int) []Recommendation {
	bookwormIdx := slices.IndexFunc(bookworms, func(bw Bookworm) bool {
		return bw.Name == bookworm
	})

	if bookwormIdx == -1 {
		// Throw error maybe, but for now return nil
		return nil
	}

	readBooks := bookworms[bookwormIdx].Books
	var recommendations map[Book]float64

	for idx, bookworm := range bookworms {
		if idx == bookwormIdx {
			continue
		}

		intersection, unreadBooks := bookIntersection(bookworm.Books, readBooks)

		var similarity float64
		similarity = float64(len(intersection))

		if similarity == 0 {
			continue
		}

		// Use logs to compress similarities incase of variation
		similarity = math.Log(similarity) + 1

		for _, ub := range unreadBooks {
			recommendations[ub] += similarity
		}
	}

	return nil
}
