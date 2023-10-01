package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	t.Run("Commom race", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)

		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expectedResponse := fastURL

		response := Race(slowURL, fastURL)

		assert(t, response, expectedResponse)
	})

	t.Run("Simultaneous race", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)

		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expectedResponse := fastURL

		response, _ := RaceButSimultaneously(slowURL, fastURL)

		assert(t, response, expectedResponse)

	})

	t.Run("Returns an error if a server doesn't respond within 10s", func(t *testing.T) {

		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		url := server.URL

		timeout := 20 * time.Millisecond

		_, err := ConfigurableRace(url, url, timeout)

		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func assert(t *testing.T, response string, expected string) {
	if response != expected {
		t.Errorf("Got %q, want %q", response, expected)
	}
}
