package pocketlog_test

import (
	pocketlog "learngo-pockets/logger"
	"testing"
)

const (
	debugMessage = "Debug message."
	infoMessage  = "Info message."
	errorMessage = "Error message."
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {

	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"Debugf": {
			level:    pocketlog.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
		},
		"Infof": {
			level:    pocketlog.LevelInfo,
			expected: infoMessage + "\n" + errorMessage + "\n",
		},
		"Errorf": {
			level:    pocketlog.LevelError,
			expected: errorMessage + "\n",
		},
	}

	for desc, tc := range tt {
		t.Run(desc, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)
			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}

		})
	}
}

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
