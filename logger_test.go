package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel(t *testing.T) {
	// level
	assert.Equal(t, Level(0), Emergency)
	assert.Equal(t, Level(1), Alert)
	assert.Equal(t, Level(2), Critical)
	assert.Equal(t, Level(3), Error)
	assert.Equal(t, Level(4), Warning)
	assert.Equal(t, Level(5), Notice)
	assert.Equal(t, Level(6), Info)
	assert.Equal(t, Level(7), Debug)

	// to int
	assert.Equal(t, 0, Emergency.Int())
	assert.Equal(t, 1, Alert.Int())
	assert.Equal(t, 2, Critical.Int())
	assert.Equal(t, 3, Error.Int())
	assert.Equal(t, 4, Warning.Int())
	assert.Equal(t, 5, Notice.Int())
	assert.Equal(t, 6, Info.Int())
	assert.Equal(t, 7, Debug.Int())

	// to string
	assert.Equal(t, "emergency", Emergency.String())
	assert.Equal(t, "alert", Alert.String())
	assert.Equal(t, "critical", Critical.String())
	assert.Equal(t, "error", Error.String())
	assert.Equal(t, "warning", Warning.String())
	assert.Equal(t, "notice", Notice.String())
	assert.Equal(t, "info", Info.String())
	assert.Equal(t, "debug", Debug.String())

	// from int
	assert.Equal(t, "emergency", Level(0).String())
	assert.Equal(t, "alert", Level(1).String())
	assert.Equal(t, "critical", Level(2).String())
	assert.Equal(t, "error", Level(3).String())
	assert.Equal(t, "warning", Level(4).String())
	assert.Equal(t, "notice", Level(5).String())
	assert.Equal(t, "info", Level(6).String())
	assert.Equal(t, "debug", Level(7).String())

	// error: This is the only one that fails
	assert.Equal(t, "", Level(8).String())
}

func TestNullLogger(t *testing.T) {
	n := NewNullLogger()

	n.Info("info")
}

func TestPrintLogger(t *testing.T) {
	p := NewPrintLogger()

	p.Emergency("emergency")
	p.Alert("alert")
	p.Critical("critical")
}
