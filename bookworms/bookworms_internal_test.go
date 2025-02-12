package main

import (
	"reflect"
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
