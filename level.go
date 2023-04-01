package logger

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

const (
	// EmergencyString is the emergency log level string
	EmergencyString = "emergency"
	// AlertString is the alert log level string
	AlertString = "alert"
	// CriticalString is the critical log level string
	CriticalString = "critical"
	// ErrorString is the error log level string
	ErrorString = "error"
	// WarningString is the warning log level string
	WarningString = "warning"
	// NoticeString is the notice log level string
	NoticeString = "notice"
	// InfoString is the info log level string
	InfoString = "info"
	// DebugString is the debug log level string
	DebugString = "debug"
)

// Int returns the integer representation of the log level
func (l Level) Int() int {
	return int(l)
}

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case Emergency:
		return EmergencyString
	case Alert:
		return AlertString
	case Critical:
		return CriticalString
	case Error:
		return ErrorString
	case Warning:
		return WarningString
	case Notice:
		return NoticeString
	case Info:
		return InfoString
	case Debug:
		return DebugString
	default:
		return ""
	}
}

func (l Level) UpperString() string {
	return strings.ToUpper(l.String())
}

func (l Level) LowerString() string {
	return strings.ToLower(l.String())
}

func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}

func (l *Level) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case EmergencyString:
		*l = Emergency
	case AlertString:
		*l = Alert
	case CriticalString:
		*l = Critical
	case ErrorString:
		*l = Error
	case WarningString:
		*l = Warning
	case NoticeString:
		*l = Notice
	case InfoString:
		*l = Info
	case DebugString:
		*l = Debug
	default:
		return fmt.Errorf("invalid log level: %s", s)
	}

	return nil
}
