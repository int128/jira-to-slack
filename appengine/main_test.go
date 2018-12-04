package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	s := httptest.NewServer(router())
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
}
