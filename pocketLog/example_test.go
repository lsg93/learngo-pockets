package pocketlog_test

import (
	pocketlog "learngo-pockets/logger"
)

func ExampleLogger_Debugf() {
	l := pocketlog.New(pocketlog.LevelDebug)
	l.Debugf("Debug message.")
	// Output: Debug message.
}

func ExampleLogger_Infof() {
	l := pocketlog.New(pocketlog.LevelInfo)
	l.Infof("Info message.")
	// Output: Info message.
}

func ExampleLogger_Errorf() {
	l := pocketlog.New(pocketlog.LevelError)
	l.Errorf("Error message.")
	// Output: Error message.
}
