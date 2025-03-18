package pocketlog_test

import (
	"fmt"
	pocketlog "learngo-pockets/logger"
)

func ExampleLogger_Debugf() {
	l := pocketlog.New(pocketlog.LevelDebug)
	debugMsg := l.Debugf("Debug message.")
	fmt.Println(debugMsg)
	// Output: Debug message.
}

func ExampleLogger_Infof() {
	l := pocketlog.New(pocketlog.LevelInfo)
	infoMsg := l.Infof("Info message.")
	fmt.Println(infoMsg)
	// Output: Info message.
}

func ExampleLogger_Errorf() {
	l := pocketlog.New(pocketlog.LevelError)
	errMsg := l.Errorf("Error message.")
	fmt.Println(errMsg)
	// Output: Error message.
}
