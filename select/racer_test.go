package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T)  {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
	
		slowServer.Close()
		fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)
	
		if err != nil {
			t.Fatalf("did not expect an error but one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		timeout := 20 * time.Millisecond
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}