package main

import (
	"math"
	"slices"
)

// The descriptions given for certain implementation details in the book are very terse/nonexistent
// So this might not be a perfect analogue to the function shown in the book @ Listing 3.22.
type Recommendation struct {
	Book  Book
	Score float64
}

func bookIntersection(bwSlice []Book, targetSlice []Book) []Book {
	set := make(map[Book]bool)

	for _, book := range bwSlice {
		set[book] = true
	}

	intersection := []Book{}

	for _, book := range targetSlice {
		if _, ok := set[book]; ok {
			intersection = append(intersection, book)
		}
	}

	return intersection
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
	recommendations := map[Book]float64{}

	for idx, bookworm := range bookworms {
		if idx == bookwormIdx {
			continue
		}

		intersection := bookIntersection(bookworm.Books, readBooks)

		var similarity float64
		similarity = float64(len(intersection))

		if similarity == 0 {
			continue
		}

		// Use logs to compress similarities incase of variation
		similarity = math.Log(similarity) + 1

		for _, book := range bookworm.Books {
			if slices.Contains(intersection, book) {
				continue
			}
			recommendations[book] += similarity
		}
	}

	results := []Recommendation{}
	for book, score := range recommendations {
		if len(results) < recommendationCount {
			results = append(results, Recommendation{Book: book, Score: score})
		}
	}

	return results
}
