package types

// Level type
type Level uint8

// AllLevels A constant exposing all logging levels
var AllLevels = []Level{
	DebugLevel,
	InfoLevel,
	WarnLevel,
	ErrorLevel,
	FatalLevel,
	PanicLevel,
}

const (
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel Level = iota

	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel

	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel

	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel

	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel

	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel
)
