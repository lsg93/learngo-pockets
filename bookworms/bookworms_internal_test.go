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
