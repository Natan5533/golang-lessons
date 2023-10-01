package race

import (
	"fmt"
	"net/http"
	"time"
)

func Race(firstCompetitor, secondCompetitor string) string {
	firstCompetitorDuration := measureResponseTime(firstCompetitor)
	secondCompetitorDuration := measureResponseTime(secondCompetitor)

	return checkWinner(firstCompetitorDuration, secondCompetitorDuration, firstCompetitor, secondCompetitor)
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func checkWinner(firstCompetitorDuration, secondCompetitorDuration time.Duration, firstCompetitor, secondCompetitor string) string {
	if firstCompetitorDuration < secondCompetitorDuration {
		return firstCompetitor
	}

	return secondCompetitor
}

// Another implementation

var timeout = 10 * time.Second

func RaceButSimultaneously(firstUrl, secondUrl string) (winner string, err error) {
	return ConfigurableRace(firstUrl, secondUrl, timeout)
}

func ConfigurableRace(firstUrl, secondUrl string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout):
		// timeout case
		return "", fmt.Errorf("timed out waiting for %s and %s", firstUrl, secondUrl)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
