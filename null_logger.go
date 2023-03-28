package logger

type NullLogger struct {
	Loggerable
}

var _ Logger = (*NullLogger)(nil)

// NewNullLogger returns a new NullLogger instance.
func NewNullLogger() *NullLogger {
	return &NullLogger{
		Loggerable: func(level Level, s string) {
			// noop
		},
	}
}
