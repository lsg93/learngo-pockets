package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello World!
}

func TestGreeting(t *testing.T) {
	expected := "Hello World!"
	actual := greeting()

	if expected != actual {
		t.Errorf("Expected string %q does not match %q", expected, actual)
	}

}
