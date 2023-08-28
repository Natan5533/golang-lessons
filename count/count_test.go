package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeperOperations{}

		Counter(buffer, spySleeper)

		result := buffer.String()
		expected := `3
2
1
Go!`

		assertEquals(t, result, expected)
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spyPrintSleep := &SpySleeperOperations{}
		Counter(spyPrintSleep, spyPrintSleep)
		expected := []string{
			"pause",
			"write",
			"pause",
			"write",
			"pause",
			"write",
			"pause",
			"write",
		}
		if !reflect.DeepEqual(expected, spyPrintSleep.Calls) {
			t.Errorf("expected calls %v, got %v", expected, spyPrintSleep.Calls)
		}
	})

	t.Run("Set time sleep", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := SetSleeper{sleepTime, spyTime.Pause}
		sleeper.Pause()

		if spyTime.durationPause != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationPause)
		}
	})
}

type SpySleeperOperations struct {
	Calls []string
}

func (s *SpySleeperOperations) Pause() {
	s.Calls = append(s.Calls, "pause")
}

func (s *SpySleeperOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func (c *SetSleeper) SetPause() {
}

func assertEquals(t *testing.T, result string, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
