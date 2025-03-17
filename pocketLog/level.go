package pocketlog

type Level byte

// Enum of thresholds at which messages should be logged.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)
