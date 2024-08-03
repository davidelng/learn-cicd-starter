package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Correct headers", func(t *testing.T) {
		h := http.Header{}
		h.Add("Authorization", "ApiKey abcdefghijkl")
		_, err := GetAPIKey(h)
		if err != nil {
			t.Errorf("Got an error but wasn't expecting one: %v", err)
		}
	})

	t.Run("Malformed headers", func(t *testing.T) {
		h := http.Header{}
		h.Add("Authorization", "abcdefghijkl")
		_, err := GetAPIKey(h)
		if err == nil {
			t.Errorf("Expecting an error but haven't got one")
		} else if err.Error() != "malformed authorization header" {
			t.Errorf("Malformed headers gave error: %v", err)
		}
	})

	t.Run("Missing headers", func(t *testing.T) {
		h := http.Header{}
		h.Add("Content-Type", "application/json")
		_, err := GetAPIKey(h)
		if err == nil {
			t.Errorf("Expecting an error but haven't got one")
		} else if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Errorf("Missing headers gave error: %v", err)
		}
	})
}
