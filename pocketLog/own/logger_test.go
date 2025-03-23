package pocketlog_own_test

import (
	"bytes"
	pocketlog_own "learngo-pockets/logger/own"
	"math/rand"
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

func TestLoggerOutputIsNotTrimmedWhenCharacterLimitIsNotMet(t *testing.T) {
	type testCase struct {
		msgLength int
		limit     int
	}

	testCases := map[string]testCase{
		"Log message is not trimmed if message length <= default character limit when none is provided": {msgLength: 1000, limit: -1},
		"Log message is not trimmed if message length <= provided character limit":                      {msgLength: 1500, limit: 1500},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			var buf bytes.Buffer
			opts := appendCharacterLimitOption(tc.limit, []pocketlog_own.LoggerOption{pocketlog_own.WithOutput(&buf)})
			l := pocketlog_own.New(opts...)

			wantResult := generateStringOfLength(tc.msgLength)

			l.Infof(wantResult)
			gotResult := strings.TrimSpace(buf.String())

			buf.Reset()

			if gotResult != wantResult {
				t.Errorf("Expected result and received result are different - expected %s, got %s", wantResult, gotResult)
			}
		})
	}
}

func TestLoggerOutputIsTrimmedWhenCharacterLimitIsMet(t *testing.T) {
	type testCase struct {
		msgLength int
		limit     int
	}

	testCases := map[string]testCase{
		"Log message is trimmed if message length > default character limit when none is provided": {msgLength: 1001, limit: -1},
		"Log message is trimmed if message length > provided character limit":                      {msgLength: 1501, limit: 1500},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			var buf bytes.Buffer
			opts := appendCharacterLimitOption(tc.limit, []pocketlog_own.LoggerOption{pocketlog_own.WithOutput(&buf)})
			l := pocketlog_own.New(opts...)

			var limit int
			if tc.limit < 0 {
				limit = 1000
			} else {
				limit = tc.limit
			}

			msg := generateStringOfLength(tc.msgLength)
			l.Infof(msg)
			gotResult := strings.TrimSpace(buf.String())

			buf.Reset()

			if len(gotResult) > limit {
				t.Errorf("Message printed by logger exceeds expected character limit.")
			}

			if !strings.HasSuffix(gotResult, "...") {
				t.Errorf("Message printed by logger was not truncated properly. Last 5 chars of result : %s", gotResult[len(gotResult)-5:])
			}
		})
	}

}

func appendCharacterLimitOption(limit int, opts []pocketlog_own.LoggerOption) []pocketlog_own.LoggerOption {
	if limit < 1000 {
		return opts
	}
	return append(opts, pocketlog_own.WithCharacterLimit(limit))
}

func generateStringOfLength(length int) string {
	bytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(bytes[rand.Intn(len(bytes))])
	}

	return sb.String()
}
