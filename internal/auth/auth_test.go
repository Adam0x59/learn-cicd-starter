package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// tests
	type test struct {
		name     string
		inputKey string
		wantKey  string
		wantErr  error
	}

	tests := []test{
		{name: "Valid Input", inputKey: "ApiKey abc123", wantKey: "abc123", wantErr: nil},
		{name: "Malformed Input", inputKey: "ApiKey abc 123", wantKey: "", wantErr: errors.New("malformed authorization header")},
		{name: "No Input", inputKey: "", wantKey: "", wantErr: ErrNoAuthHeaderIncluded},
	}

	for _, test := range tests {
		// arrange
		headers := http.Header{}
		headers.Set("Authorization", test.inputKey)

		// act
		key, err := GetAPIKey(headers)

		// assert
		if test.wantErr != nil {
			if err == nil || err.Error() != test.wantErr.Error() {
				t.Fatalf("expected error %v, got %v", test.wantErr, err)
			}
		} else {
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if key != test.wantKey {
				t.Errorf("expected '%s', got '%s'", test.wantKey, key)
			}
		}
	}
}
