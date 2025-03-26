package pocketlog_own

import (
	"encoding/json"
	"fmt"
)

// We use something similar to the strategy pattern to define different possible formats for the logger's output.

type LoggerMessageFormatter interface {
	Format(input string, ctx LoggerContext) string
}

type PlaintextFormatter struct{}

func (ptf *PlaintextFormatter) Format(format string, ctx LoggerContext) string {
	return format + "\n"
}

type JSONFormatter struct{}

type LoggerMessageJSON struct {
	Level   string
	Message string
}

func (jf *JSONFormatter) Format(format string, ctx LoggerContext) string {

	jsonMsg := LoggerMessageJSON{
		Level:   ctx.level.String(),
		Message: format,
	}

	jsonStr, err := json.Marshal(&jsonMsg)

	if err != nil {
		fmt.Errorf("An error occurred while marshaling JSON.")
	}

	return string(jsonStr)
}
