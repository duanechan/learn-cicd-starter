package auth

import (
	"net/http"
	"testing"
)

var testCases = []struct {
	input    http.Header
	expected string
}{
	{
		input:    http.Header{"Authorization": []string{"ApiKey 123-Imaginary-Api-Key-123"}},
		expected: "123-Imaginary-Api-Key-123",
	},
	{
		input:    http.Header{"Content-Type": []string{"application/json"}},
		expected: "",
	},
	{
		input:    http.Header{"Authorization": []string{"ApiKey"}},
		expected: "",
	},
}

func TestGetAPIKey(t *testing.T) {
	for _, test := range testCases {
		actual, err := GetAPIKey(test.input)
		if test.expected != "" && err != nil {
			t.Errorf("expected no err, got %s", err)
		}
		if test.expected != actual {
			t.Errorf("expected result %s, got %s", test.expected, actual)
		}
	}
}
