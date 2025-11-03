package auth

import (
	"errors"
	"net/http"
	"testing"
)

type result struct {
	result string
	err    error
}

var testCases = []struct {
	input    http.Header
	expected result
}{
	{
		input:    http.Header{"Authorization": []string{"ApiKey 123-Imaginary-Api-Key-123"}},
		expected: result{"123-Imaginary-Api-Key-123", nil},
	},
	{
		input:    http.Header{"Content-Type": []string{"application/json"}},
		expected: result{"", ErrNoAuthHeaderIncluded},
	},
	{
		input:    http.Header{"Authorization": []string{"ApiKey"}},
		expected: result{"", errors.New("malformed authorization header")},
	},
}

func TestGetAPIKey(t *testing.T) {
	for _, test := range testCases {
		apiKey, err := GetAPIKey(test.input)
		actual := result{result: apiKey, err: err}
		if test.expected != actual {
			t.Errorf("expected %s, got %s", test.expected.result, actual.result)
		}
	}
}
