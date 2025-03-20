package pocketlog_own_test

import (
	pocketlog_own "learngo-pockets/logger"
	"os"
)

func ExampleLogger_Debugf() {
	l := pocketlog_own.New(os.Stdout, pocketlog_own.LevelDebug)
	l.Debugf("Debug message.")
	// Output: Debug message.
}

func ExampleLogger_Infof() {
	l := pocketlog_own.New(os.Stdout, pocketlog_own.LevelInfo)
	l.Infof("Info message.")
	// Output: Info message.
}

func ExampleLogger_Errorf() {
	l := pocketlog_own.New(os.Stdout, pocketlog_own.LevelError)
	l.Errorf("Error message.")
	// Output: Error message.
}
