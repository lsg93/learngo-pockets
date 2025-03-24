package pocketlog_own_test

import (
	"bytes"
	"encoding/json"
	pocketlog_own "learngo-pockets/logger/own"
	"strings"
	"testing"
)

var message string = "Logger output."

func TestLoggerWithPlaintextFormatter(t *testing.T) {

	var buf bytes.Buffer
	wantResult := message
	// This test also tests that the default formatter used by the logger is the PlaintextFormatter.
	l := pocketlog_own.New(pocketlog_own.WithOutput(&buf))

	l.Infof(message)

	gotResult := strings.TrimSpace(buf.String())

	if gotResult != wantResult {
		t.Errorf("Expected result differs from received result : Want %s, got %s", wantResult, gotResult)
	}
}

func TestLoggerWithJSONFormatter(t *testing.T) {
	var buf bytes.Buffer

	type testLoggerJSONFormat struct {
		level     pocketlog_own.LoggerLevel
		message   string
		timestamp string // For now...
	}

	wantResult := json.Marshal()
	// This test also tests that the default formatter used by the logger is the PlaintextFormatter.
	l := pocketlog_own.New(pocketlog_own.WithOutput(&buf))

	l.Infof(message)

	gotResult := strings.TrimSpace(buf.String())

	if gotResult != wantResult {
		t.Errorf("Expected result differs from received result : Want %s, got %s", wantResult, gotResult)
	}
}
