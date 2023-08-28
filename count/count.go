package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pause()
}

func Counter(result io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Pause()
		fmt.Fprintln(result, i)
	}
	sleeper.Pause()
	fmt.Fprint(result, "Go!")

}

func main() {
	sleeper := &SetSleeper{5 * time.Second, time.Sleep}
	Counter(os.Stdout, sleeper)
}

type SetSleeper struct {
	duration time.Duration
	pause    func(time.Duration)
}

type SpyTime struct {
	durationPause time.Duration
}

func (s *SpyTime) Pause(duration time.Duration) {
	s.durationPause = duration
}

func (s *SetSleeper) Pause() {
	s.pause(s.duration)
}
