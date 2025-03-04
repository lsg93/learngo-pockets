package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestLoadBoowkorms(t *testing.T) {

	type testCase struct {
		fp      string
		want    []Bookworm
		wantErr bool
	}

	testCases := map[string]testCase{
		"file exists": {
			fp: "testdata/test_valid_bookworms.json",
			want: []Bookworm{
				{
					Name: "Jane",
					Books: []Book{
						{Author: "Agatha Christie", Name: "And Then There Were None"},
						{Author: "Stephen King", Name: "The Shining"},
					},
				},
			},
			wantErr: false,
		},
		"invalid filepath": {
			fp:      "testdata/non_existent_file.json",
			want:    nil,
			wantErr: true,
		},
		"malformed JSON": {
			fp:      "testdata/test_malformed_bookworms.json",
			want:    nil,
			wantErr: true,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			got, err := loadBookworms(tc.fp)

			if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Values for loaded bookworms and expected bookworms do not match. Expected %q , got %q", tc.want, got)
			}

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error, but did not receive one")
				}
			}
		})
	}
}

func TestGetBookCounts(t *testing.T) {

	type testCase struct {
		slice      []Bookworm
		wantResult BookCount
	}

	testCases := map[string]testCase{
		"handles multiple books from multiple users": {
			slice: []Bookworm{
				{Name: "Lawrence", Books: []Book{
					{Name: "Tiny CSS Projects", Author: "Martine Dowden"},
					{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"},
				}},
				{Name: "Samuel", Books: []Book{
					{Name: "Tiny CSS Projects", Author: "Martine Dowden"},
					{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"},
				}},
			},
			wantResult: BookCount{
				{Name: "Tiny CSS Projects", Author: "Martine Dowden"}:                   2,
				{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"}: 2,
			},
		},
		"Handles multiple books from one user if duplicates exist": {
			slice: []Bookworm{
				{Name: "Lawrence", Books: []Book{
					{Name: "Tiny CSS Projects", Author: "Martine Dowden"},
					{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"},
					{Name: "Tiny CSS Projects", Author: "Martine Dowden"},
					{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"},
				}},
			},
			wantResult: BookCount{
				{Name: "Tiny CSS Projects", Author: "Martine Dowden"}:                   2,
				{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"}: 2,
			},
		},
		"Handles bookworm with no books": {
			slice: []Bookworm{
				{Name: "Lawrence", Books: []Book{
					{Name: "Tiny CSS Projects", Author: "Martine Dowden"},
					{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"},
				}},
				{Name: "Samuel", Books: []Book{}},
			},
			wantResult: BookCount{
				{Name: "Tiny CSS Projects", Author: "Martine Dowden"}:                   1,
				{Name: "Learn Go With Pocket-Sized Projects", Author: "Alienor Latour"}: 1,
			},
		},
		"Handles case where no bookworms are given": {
			slice:      []Bookworm{},
			wantResult: BookCount{},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult := getBookCounts(tc.slice)

			if !checkBookMapsAreEqual(gotResult, tc.wantResult) {
				t.Errorf("Expected %q, got %q", gotResult, tc.wantResult)
			}

		})
	}
}

func TestFindCommonBooks(t *testing.T) {

	type testCase struct {
		slice      []Bookworm
		wantResult []Book
	}

	book1 := Book{Author: "Author1", Name: "Book1"}
	book2 := Book{Author: "Author2", Name: "Book2"}
	book3 := Book{Author: "Author3", Name: "Book3"}
	book4 := Book{Author: "Author4", Name: "Book4"}
	book5 := Book{Author: "Author5", Name: "Book5"}
	book6 := Book{Author: "Author6", Name: "Book6"}
	book7 := Book{Author: "Author7", Name: "Book7"}

	testCases := map[string]testCase{
		"Every user same books": {
			slice: []Bookworm{
				{Name: "User1", Books: []Book{book1, book2}},
				{Name: "User2", Books: []Book{book1, book2}},
				{Name: "User3", Books: []Book{book1, book2}},
			},
			wantResult: []Book{book1, book2},
		},
		"Multiple common books between users": {
			slice: []Bookworm{
				{Name: "User4", Books: []Book{book1, book2, book3}},
				{Name: "User5", Books: []Book{book1, book3, book4}},
				{Name: "User6", Books: []Book{book2, book3}},
			},
			wantResult: []Book{book1, book2, book3},
		},
		"No common books between users": {
			slice: []Bookworm{
				{Name: "User7", Books: []Book{book5}},
				{Name: "User8", Books: []Book{book6}},
				{Name: "User9", Books: []Book{book7}},
			},
			wantResult: nil,
		},
		"Some users with no books": {
			slice: []Bookworm{
				{Name: "User7", Books: []Book{book5}},
				{Name: "User8", Books: []Book{book6}},
				{Name: "User9", Books: []Book{book7}},
			},
			wantResult: nil,
		},
		"No users with any books": {
			slice: []Bookworm{
				{Name: "User12", Books: []Book{book1}},
				{Name: "User13", Books: []Book{}},
				{Name: "User14", Books: []Book{book4}},
			},
			wantResult: nil,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult := findCommonBooks(tc.slice)
			if !slices.Equal(tc.wantResult, gotResult) {
				t.Errorf("Common books are not returned as expected. Expected %v got %v", tc.wantResult, gotResult)
			}
		})
	}
}

func checkBookMapsAreEqual(mapA BookCount, mapB BookCount) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for key, valueA := range mapA {
		valueB, ok := mapB[key]
		if !ok || valueA != valueB {
			return false
		}
	}
	return true
}
