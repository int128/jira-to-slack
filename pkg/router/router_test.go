package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	s := httptest.NewServer(New())
	defer s.Close()

	t.Run("index", func(t *testing.T) {
		resp, err := http.Get(s.URL + "/?webhook=dummy")
		if err != nil {
			t.Fatalf("Could not get %s: %s", s.URL, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("StatusCode wants 400 but %d", resp.StatusCode)
		}
	})
	t.Run("healthz", func(t *testing.T) {
		resp, err := http.Get(s.URL + "/healthz")
		if err != nil {
			t.Fatalf("Could not get %s: %s", s.URL+"/healthz", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("StatusCode wants 200 but %d", resp.StatusCode)
		}
	})
}
