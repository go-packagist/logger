package logger

type Level int

// see https://www.rfc-editor.org/rfc/rfc5424.html
const (
	// Emergency is the emergency log level
	Emergency Level = iota
	// Alert is the alert log level
	Alert
	// Critical is the critical log level
	Critical
	// Error is the error log level
	Error
	// Warning is the warning log level
	Warning
	// Notice is the notice log level
	Notice
	// Info is the info log level
	Info
	// Debug is the debug log level
	Debug
)

// Int returns the integer representation of the log level
func (l Level) Int() int {
	return int(l)
}

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case Emergency:
		return "emergency"
	case Alert:
		return "alert"
	case Critical:
		return "critical"
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Notice:
		return "notice"
	case Info:
		return "info"
	case Debug:
		return "debug"
	default:
		return ""
	}
}
