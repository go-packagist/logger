package logger

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	Level Level `json:"level"`
}

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

	// upper string
	assert.Equal(t, "EMERGENCY", Emergency.UpperString())

	// lower string
	assert.Equal(t, "emergency", Emergency.LowerString())
}

func TestLevel_Marshal(t *testing.T) {
	t1 := test{Level: Emergency}
	t2 := test{Level: Alert}
	t3 := test{Level: Critical}
	t4 := test{Level: Error}
	t5 := test{Level: Warning}
	t6 := test{Level: Notice}
	t7 := test{Level: Info}
	t8 := test{Level: Debug}

	j1, _ := json.Marshal(t1)
	j2, _ := json.Marshal(t2)
	j3, _ := json.Marshal(t3)
	j4, _ := json.Marshal(t4)
	j5, _ := json.Marshal(t5)
	j6, _ := json.Marshal(t6)
	j7, _ := json.Marshal(t7)
	j8, _ := json.Marshal(t8)

	assert.Equal(t, `{"level":"emergency"}`, string(j1))
	assert.Equal(t, `{"level":"alert"}`, string(j2))
	assert.Equal(t, `{"level":"critical"}`, string(j3))
	assert.Equal(t, `{"level":"error"}`, string(j4))
	assert.Equal(t, `{"level":"warning"}`, string(j5))
	assert.Equal(t, `{"level":"notice"}`, string(j6))
	assert.Equal(t, `{"level":"info"}`, string(j7))
	assert.Equal(t, `{"level":"debug"}`, string(j8))
}

func TestLevel_Unmarshal(t *testing.T) {
	s1 := `{"level":"emergency"}`
	s2 := `{"level":"alert"}`
	s3 := `{"level":"critical"}`
	s4 := `{"level":"error"}`
	s5 := `{"level":"warning"}`
	s6 := `{"level":"notice"}`
	s7 := `{"level":"info"}`
	s8 := `{"level":"debug"}`
	s9 := `{"level":"test"}`
	s10 := `{"level":8}`

	var t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 test
	_ = json.Unmarshal([]byte(s1), &t1)
	_ = json.Unmarshal([]byte(s2), &t2)
	_ = json.Unmarshal([]byte(s3), &t3)
	_ = json.Unmarshal([]byte(s4), &t4)
	_ = json.Unmarshal([]byte(s5), &t5)
	_ = json.Unmarshal([]byte(s6), &t6)
	_ = json.Unmarshal([]byte(s7), &t7)
	_ = json.Unmarshal([]byte(s8), &t8)
	err1 := json.Unmarshal([]byte(s9), &t9)
	err2 := json.Unmarshal([]byte(s10), &t10)

	assert.Equal(t, Emergency, t1.Level)
	assert.Equal(t, Alert, t2.Level)
	assert.Equal(t, Critical, t3.Level)
	assert.Equal(t, Error, t4.Level)
	assert.Equal(t, Warning, t5.Level)
	assert.Equal(t, Notice, t6.Level)
	assert.Equal(t, Info, t7.Level)
	assert.Equal(t, Debug, t8.Level)
	assert.Error(t, err1)
	assert.Equal(t, "invalid log level: test", err1.Error())
	assert.Error(t, err2)
}

func TestNullLogger(t *testing.T) {
	n := NewNullLogger()

	n.Info("info")
}

func TestPrintLogger(t *testing.T) {
	p := NewPrintLogger()

	p.Emergencyf("emergency %s", "test")
	p.Alertf("alert %s", "test")
	p.Criticalf("critical %s", "test")
	p.Errorf("error %s", "test")
	p.Warningf("warning %s", "test")
	p.Noticef("notice %s", "test")
	p.Infof("info %s", "test")
	p.Debugf("debug %s", "test")

	p.Emergency("emergency")
	p.Alert("alert")
	p.Critical("critical")
	p.Error("error")
	p.Warning("warning")
	p.Notice("notice")
	p.Info("info")
	p.Debug("debug")

	p.Log(Emergency, "emergency")
}
