package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		errorMsg  string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "return lowercase",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/PATH/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://OTHER.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute path",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://other.com/path/one"},
		},
		{
			name:     "relative path",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href=":\\invalidURL">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: nil,
			errorMsg: "couldn't parse href",
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: nil,
		},
		{
			name:     "bad html",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html body>
				<a href="path/one">
					<span>Boot.dev</span>
				</a>
			</html body>
		`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorMsg) {
				t.Errorf("TEST %v - '%s' FAILED: unexpected error: %v\n", i, tc.name, err)
				return
			} else if err != nil && tc.errorMsg == "" {
				t.Errorf("TEST %v - '%s' FAILED: unexpected error: %v\n", i, tc.name, err)
				return
			} else if err != nil && tc.errorMsg != "" {
				t.Errorf("TEST %v - '%s' FAILED: expected error containing '%v', got '%v':\n", i, tc.name, tc.errorMsg, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("TEST %v - '%s' FAILED: got '%v', expected '%v'\n", i, tc.name, actual, tc.expected)
				return
			}
		})
	}
}
