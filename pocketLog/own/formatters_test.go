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
		Level   string
		Message string
	}

	jsonStr, _ := json.Marshal(testLoggerJSONFormat{Level: "Info", Message: message})
	wantResult := string(jsonStr)
	l := pocketlog_own.New(pocketlog_own.WithFormatter(
		&pocketlog_own.JSONFormatter{}),
		pocketlog_own.WithOutput(&buf),
	)

	l.Infof(message)

	gotResult := buf.String()

	if gotResult != wantResult {
		t.Errorf("Expected result differs from received result : Want %s, got %s", wantResult, gotResult)
	}
}
