package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		errorMsg string
	}{
		{
			name:     "removed scheme",
			input:    "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "removed trailing slash",
			input:    "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove upper case",
			input:    "https://BLOG.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "invalid url",
			input:    "BLOG.boot.dev/path/",
			expected: "",
			errorMsg: "invalid url: no http protocol",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.input)
			if err != nil && !strings.Contains(err.Error(), tc.errorMsg) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v\n", i, tc.name, err)
				return
			} else if err != nil && tc.errorMsg == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v\n", i, tc.name, err)
				return
			} else if err == nil && tc.errorMsg != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error conaining '%v', got none\n", i, tc.name, tc.errorMsg)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
