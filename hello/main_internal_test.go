package main

import (
	"fmt"
	"testing"
)

// Rather than creating another map just for testing purposes, iterate over the map from main.go.
// This way tests are always in sync with the data that's actually in use.

// Similarly to the main.go, I've tested for errors (probably badly) even though the book ignores it at this early stage.
func TestGreeting(t *testing.T) {

	t.Run("Works with country codes from phrasebook", func(t *testing.T) {
		for code, phrase := range Phrasebook {
			got, _ := greeting(code)
			want := phrase
			if got != want {
				t.Errorf("Expected string %q does not match %q", got, want)
			}

		}
	})

	t.Run("Errors as expected when no country code is given", func(t *testing.T) {
		_, gotErrMsg := greeting("")
		if gotErrMsg == nil {
			t.Errorf("Expected error, but did not get one.")
		}
		wantErrMsg := "no country code provided."
		if gotErrMsg.Error() != wantErrMsg {
			t.Errorf("Expected error mesgosage %q, got %q", wantErrMsg, gotErrMsg)
		}
	})

	t.Run("Errors as expected when country code is not found in phrasebook", func(t *testing.T) {
		_, gotErrMsg := greeting("abc")
		if gotErrMsg == nil {
			t.Errorf("country code %q is not supported", "abc")
		}
		wantErrMsg := fmt.Sprintf("country code %q is not supported", "abc")
		if gotErrMsg.Error() != wantErrMsg {
			t.Errorf("Expected error message %q, got %q", wantErrMsg, gotErrMsg)
		}
	})

}
