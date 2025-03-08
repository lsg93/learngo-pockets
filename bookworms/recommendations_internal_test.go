package main

import (
	"slices"
	"testing"
)

var book1 = Book{Author: "Author1", Name: "Book1"}
var book2 = Book{Author: "Author2", Name: "Book2"}
var book3 = Book{Author: "Author3", Name: "Book3"}
var book4 = Book{Author: "Author4", Name: "Book4"}
var book5 = Book{Author: "Author5", Name: "Book5"}
var book6 = Book{Author: "Author6", Name: "Book6"}
var book7 = Book{Author: "Author7", Name: "Book7"}

func TestBookIntersection(t *testing.T) {
	type TestCase struct {
		slice1     []Book
		slice2     []Book
		wantResult []Book
	}

	testCases := map[string]TestCase{
		"All shared books": {
			slice1:     []Book{book1, book2, book3},
			slice2:     []Book{book1, book2, book3},
			wantResult: []Book{book1, book2, book3},
		},
		"Some shared books": {
			slice1:     []Book{book1, book2, book3},
			slice2:     []Book{book2, book3, book4},
			wantResult: []Book{book2, book3},
		},
		"No shared books": {
			slice1:     []Book{book1, book2, book3},
			slice2:     []Book{book4, book5, book6},
			wantResult: []Book{},
		},
		"No books given": {
			slice1:     []Book{},
			slice2:     []Book{},
			wantResult: []Book{},
		},
	}

	//TODO fix this test...
	// Need to make wantResult have 2 values basically.

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult, gotResult2 := bookIntersection(tc.slice1, tc.slice2)
			if !slices.Equal(gotResult, gotResult2) {
				t.Errorf("Results of intersection were different to what was expected - %v expected, %v returned", tc.wantResult, gotResult)
			}
		})
	}
}

func TestRecommendations(t *testing.T) {
	type TestCase struct {
		bookworms      []Bookworm
		chosenBookworm string
		wantResult     []Recommendation
	}

	testCases := map[string]TestCase{
		"bookworms with some books in common": {
			chosenBookworm: "User1",
			bookworms: []Bookworm{
				{Name: "User1", Books: []Book{book1, book2, book3}},
				{Name: "User2", Books: []Book{book1, book3, book4}},
				{Name: "User3", Books: []Book{book2, book3}},
			},
			wantResult: []Recommendation{{Book: book1, score: 1}, {Book: book4, score: 1}},
		},
		// "bookworms with no books in common": {
		// 	chosenBookworm: {},
		// 	bookworms:      []Bookworm{},
		// 	wantResult:   []Recommendation{},
		// },
		// "bookworms with identical books": {
		// 	chosenBookworm: {},
		// 	bookworms:      []Bookworm{},
		// 	wantResult:   []Recomendation{},
		// },
		// "No bookworms given": {
		// 	chosenBookworm: {},
		// 	bookworms:      []Bookworm{},
		// 	wantResult:   []Recomendation{},
		// },
		// "Reader given does not exist in range of bookworms": {
		// }
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult := recommendBooks(tc.chosenBookworm, tc.bookworms, len(tc.wantResult))
			if !slices.Equal(tc.wantResult, gotResult) {
				t.Errorf("Recommendations given are not those that were expected : Expected %v, received %v", tc.wantResult, gotResult)
			}
		})
	}

}
