package pocketlog_own

// TODO - come back when you learn how to test against io.Writer

// func TestLogLevelIsComputedAutomatically(t *testing.T) {

// 	type TestCase struct {
// 		logFunc    func(format string) string
// 		wantResult bool
// 	}

// 	testCases := map[string]TestCase{
// 		"Logs debug messages when threshold is set at debug level":         {logFunc: func(s string) string { return New(LevelDebug).Debugf(s) }, wantResult: true},
// 		"Logs info messages when threshold is set at debug level":          {logFunc: func(s string) string { return New(LevelDebug).Infof(s) }, wantResult: true},
// 		"Logs error messages when threshold is set at debug level":         {logFunc: func(s string) string { return New(LevelDebug).Errorf(s) }, wantResult: true},
// 		"Does not log debug messages when threshold is set at info level":  {logFunc: func(s string) string { return New(LevelInfo).Debugf(s) }, wantResult: false},
// 		"Logs info messages when threshold is set at info level":           {logFunc: func(s string) string { return New(LevelInfo).Infof(s) }, wantResult: true},
// 		"Logs error messages when threshold is set at info level":          {logFunc: func(s string) string { return New(LevelInfo).Errorf(s) }, wantResult: true},
// 		"Does not log debug messages when threshold is set at error level": {logFunc: func(s string) string { return New(LevelError).Debugf(s) }, wantResult: false},
// 		"Does not log info messages when threshold is set at error level":  {logFunc: func(s string) string { return New(LevelError).Infof(s) }, wantResult: false},
// 		"Logs error messages when threshold is set at error level":         {logFunc: func(s string) string { return New(LevelError).Errorf(s) }, wantResult: true},
// 	}

// 	message := "Test message"

// 	for desc, tc := range testCases {
// 		t.Run(desc, func(t *testing.T) {
// 			gotResult := len(tc.logFunc(message)) != 0
// 			if gotResult != tc.wantResult {
// 				t.Errorf("An error has occured - message was logged but it was not supposed to be for given error level.")
// 			}
// 		})
// 	}
// }
