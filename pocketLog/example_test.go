package pocketlog_test

import (
	pocketlog "learngo-pockets/logger"
	"os"
)

func ExampleLogger_Debugf() {
	l := pocketlog.New(os.Stdout, pocketlog.LevelDebug)
	l.Debugf("Debug message.")
	// Output: Debug message.
}

func ExampleLogger_Infof() {
	l := pocketlog.New(os.Stdout, pocketlog.LevelInfo)
	l.Infof("Info message.")
	// Output: Info message.
}

func ExampleLogger_Errorf() {
	l := pocketlog.New(os.Stdout, pocketlog.LevelError)
	l.Errorf("Error message.")
	// Output: Error message.
}
