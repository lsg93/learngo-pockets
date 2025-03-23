package pocketlog_own_test

import (
	"bytes"
	pocketlog_own "learngo-pockets/logger/own"
	"strings"
	"testing"
)

const (
	debugMsg = "Debug message."
	infoMsg  = "Info message."
	errMsg   = "Error message."
)

type LoggerMethod string

var Debugf LoggerMethod = "Debugf"
var Infof LoggerMethod = "Infof"
var Errorf LoggerMethod = "Errorf"

type LoggerOutput map[LoggerMethod]string

// This test is awful and trying to be too clever. Would be better if broken down into separate tests for each method.
// Trying too hard to ensure all methods are tested, but Go doesn't seem to work like that.
func TestLoggerMethodsRespectThresholds(t *testing.T) {

	type testCase struct {
		threshold  pocketlog_own.LoggerLevel
		wantResult LoggerOutput
	}

	testCases := map[string]testCase{
		"Logs debug/info/error messages when threshold is set at debug level": {threshold: pocketlog_own.LevelDebug, wantResult: LoggerOutput{Debugf: debugMsg, Infof: infoMsg, Errorf: errMsg}},
		"Only logs info/error messages when threshold is set at info level":   {threshold: pocketlog_own.LevelInfo, wantResult: LoggerOutput{Infof: infoMsg, Errorf: errMsg}},
		"Only logs error messages when threshold is set at error level":       {threshold: pocketlog_own.LevelError, wantResult: LoggerOutput{Errorf: errMsg}},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			// Create writer to use as output
			tlw := testLoggerWriter{contents: make(LoggerOutput)}

			// Create a logger at threshold
			l := pocketlog_own.New(
				pocketlog_own.WithOutput(&tlw),
				pocketlog_own.WithThreshold(tc.threshold),
			)

			// Log messages - before logging, set the method in our testLoggerWriter struct so the message is added to the map.
			tlw.setMethod("Debugf")
			l.Debugf(debugMsg)
			tlw.setMethod("Infof")
			l.Infof(infoMsg)
			tlw.setMethod("Errorf")
			l.Errorf(errMsg)

			// Take data from buffer and parse into a map which we can assert against.
			gotResult := tlw.contents

			if !compareLoggerOutput(gotResult, tc.wantResult) {
				t.Errorf("An error has occured - there is a difference between the wanted output and the expected output.")
			}
		})
	}

}

type testLoggerWriter struct {
	contents LoggerOutput
	method   LoggerMethod
}

func (tlw *testLoggerWriter) setMethod(method LoggerMethod) {
	tlw.method = method
}

func (tlw *testLoggerWriter) Write(p []byte) (n int, err error) {
	// Remove newlines for testing purposes.
	tlw.contents[tlw.method] = strings.TrimSpace(string(p))
	return len(p), nil
}

func compareLoggerOutput(map1 LoggerOutput, map2 LoggerOutput) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, map1Value := range map1 {
		if map2Value, ok := map2[key]; !ok || map1Value != map2Value {
			return false
		}
	}

	return true
}

func TestLoggerMethodsDoNotLogLogLevelByDefault(t *testing.T) {
	var buf bytes.Buffer
	l := pocketlog_own.New(
		pocketlog_own.WithOutput(&buf),
		pocketlog_own.WithThreshold(pocketlog_own.LevelDebug),
	)

	type testCase struct {
		logMethod  func(format string, args ...any)
		msg        string
		wantResult string
	}

	testCases := map[string]testCase{
		"Debugf method does not show level when logging messages": {
			logMethod: l.Debugf, msg: debugMsg, wantResult: debugMsg,
		},
		"Infof method does not show level when logging messages": {
			logMethod: l.Infof, msg: infoMsg, wantResult: infoMsg,
		},
		"Errorf method does not show level when logging messages": {
			logMethod: l.Errorf, msg: errMsg, wantResult: errMsg,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.logMethod(tc.msg)
			// Trim newlines
			gotResult := strings.TrimSpace(buf.String())

			buf.Reset()

			if gotResult != tc.wantResult {
				t.Errorf("Expected result %s is different to received result %s", gotResult, tc.wantResult)
			}
		})
	}

}

func TestLoggerMethodsLogLogLevel(t *testing.T) {

	var buf bytes.Buffer
	l := pocketlog_own.New(
		pocketlog_own.WithLogLevel(),
		pocketlog_own.WithOutput(&buf),
		pocketlog_own.WithThreshold(pocketlog_own.LevelDebug),
	)

	type testCase struct {
		logMethod  func(format string, args ...any)
		msg        string
		wantResult string
	}

	testCases := map[string]testCase{
		"Debugf method shows log level when logging messages": {
			logMethod: l.Debugf, msg: debugMsg, wantResult: "[DEBUG]: " + debugMsg,
		},
		"Infof method shows log level when logging messages": {
			logMethod: l.Infof, msg: infoMsg, wantResult: "[INFO]: " + infoMsg,
		},
		"Errorf method shows log level when logging messages": {
			logMethod: l.Errorf, msg: errMsg, wantResult: "[ERROR]: " + errMsg,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			tc.logMethod(tc.msg)
			// Trim newlines
			gotResult := strings.TrimSpace(buf.String())

			buf.Reset()

			if gotResult != tc.wantResult {
				t.Errorf("Expected result %s is different to received result %s", gotResult, tc.wantResult)
			}
		})
	}
}
