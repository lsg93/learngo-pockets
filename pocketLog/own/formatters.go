package pocketlog_own

// We use something similar to the strategy pattern to define different possible formats for the logger's output

type LoggerMessageFormatter interface {
	Format(input string) string
}

type PlaintextFormatter struct{}

func (ptf *PlaintextFormatter) Format(s string) string {
	return s
}

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(s string) string {
	// Take a string and encode it into
	/**
	{
		level : LoggerLevel
		message : ?
		timestamp : Time?
	}
	*/
}
