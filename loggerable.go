package logger

import "fmt"

type Loggerable func(Level, string)

func (l Loggerable) Emergencyf(format string, args ...interface{}) {
	l.Emergency(fmt.Sprintf(format, args...))
}

func (l Loggerable) Alertf(format string, args ...interface{}) {
	l.Alert(fmt.Sprintf(format, args...))
}

func (l Loggerable) Criticalf(format string, args ...interface{}) {
	l.Critical(fmt.Sprintf(format, args...))
}

func (l Loggerable) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l Loggerable) Warningf(format string, args ...interface{}) {
	l.Warning(fmt.Sprintf(format, args...))
}

func (l Loggerable) Noticef(format string, args ...interface{}) {
	l.Notice(fmt.Sprintf(format, args...))
}

func (l Loggerable) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l Loggerable) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l Loggerable) Emergency(message string) {
	l.Log(Emergency, message)
}

func (l Loggerable) Alert(message string) {
	l.Log(Alert, message)
}

func (l Loggerable) Critical(message string) {
	l.Log(Critical, message)
}

func (l Loggerable) Error(message string) {
	l.Log(Error, message)
}

func (l Loggerable) Warning(message string) {
	l.Log(Warning, message)
}

func (l Loggerable) Notice(message string) {
	l.Log(Notice, message)
}

func (l Loggerable) Info(message string) {
	l.Log(Info, message)
}

func (l Loggerable) Debug(message string) {
	l.Log(Debug, message)
}

func (l Loggerable) Log(level Level, message string) {
	l(level, message)
}
