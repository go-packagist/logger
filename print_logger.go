package logger

import (
	"fmt"
	"time"
)

type PrintLogger struct {
	Loggerable
}

var _ Logger = (*PrintLogger)(nil)

// NewPrintLogger returns a new PrintLogger instance.
func NewPrintLogger() *PrintLogger {
	return &PrintLogger{
		Loggerable: func(level Level, s string) {
			println(fmt.Sprintf("%s %s: %s", time.Now().Format(time.DateTime), level, s))
		},
	}
}
