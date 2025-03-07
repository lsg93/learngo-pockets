package main

import (
	"flag"
	"fmt"
	"strings"
)

type language string

/*
	The example code in the book has the Phrasebook variable lowercased and not exported.
	The test has basically the exact same map and iterates over it to generate tests
	Even though this is supposed to be a simple project,
	Declaring two maps means that both things would need to be updated if new languages were ever added.
	Opting to export the variable means it can be reused in the test file.
*/

const (
	ErrNoCountryCode          = "no country code provided."
	ErrUnsupportedCountryCode = "country code %q is not supported"
)

var Phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ﻮ ﻠ ﯿ ﮨ",
	"vi": "Xin chào Thế Giới",
}

// Opted to try and handle errors in functions below even though the book says that comes in later chapters.
func main() {
	var code string
	flag.StringVar(&code, "code", "en", `The two-character ISO 639-1 language code for the language you'd like to be greeted in, e.g. "en"`)
	flag.Parse()
	msg, err := greeting(language(code))
	if err != nil {
		fmt.Printf("Error : %q", err)
	}
	fmt.Println(msg)
}

func greeting(l language) (string, error) {
	msg, ok := Phrasebook[l]
	if !ok {
		var errMsg string
		if strings.TrimSpace(string(l)) == "" {
			errMsg = ErrNoCountryCode
		} else {
			errMsg = fmt.Sprintf(ErrUnsupportedCountryCode, l)
		}
		return "", fmt.Errorf(errMsg)
	}
	return msg, nil
}
