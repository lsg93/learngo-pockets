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

var loggerMethods = []LoggerMethod{Debugf, Infof, Errorf}

type LoggerOutput map[LoggerMethod]string

func TestLoggerMethodsRespectThresholds(t *testing.T) {

	type testCase struct {
		threshold  pocketlog_own.LoggerLevel
		wantResult LoggerOutput
	}

	testCases := map[string]testCase{
		"Logs debug/info/error messages when threshold is set at debug level": {threshold: pocketlog_own.LevelDebug, wantResult: LoggerOutput{Debugf: debugMsg, Infof: infoMsg, Errorf: errMsg}},
		"Only logs info/error messages when threshold is set at info level":   {threshold: pocketlog_own.LevelInfo, wantResult: LoggerOutput{Debugf: "", Infof: infoMsg, Errorf: errMsg}},
		"Only logs error messages when threshold is set at error level":       {threshold: pocketlog_own.LevelInfo, wantResult: LoggerOutput{Debugf: "", Infof: "", Errorf: errMsg}},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			// Create buffer to use as output
			var buf bytes.Buffer

			// Create a logger at threshold
			l := pocketlog_own.New(
				pocketlog_own.WithOutput(&buf),
				pocketlog_own.WithThreshold(tc.threshold),
			)

			// Log messages
			l.Debugf(debugMsg)
			l.Infof(infoMsg)
			l.Errorf(errMsg)

			// Take data from buffer and parse into a map which we can assert against.
			gotResult := marshalLoggerOutput(&buf)

			if !compareLoggerOutput(gotResult, tc.wantResult) {
				t.Errorf("An error has occured - there is a difference between the wanted output and the expected output.")
			}
		})
	}

}

func marshalLoggerOutput(b *bytes.Buffer) LoggerOutput {
	o := make(LoggerOutput)
	s := strings.Split(b.String(), "\n")

	if len(s) != len(loggerMethods) {
		// Not sure what to do here, but there should be an error.
		// This way if new logger methods are added without the tests being updated, they wil fail,
		// In other languages I'd throw/return an error, but not sure how to do that in Go just yet.
	}

	for idx, method := range loggerMethods {
		o[method] = s[idx]
	}

	return o
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
