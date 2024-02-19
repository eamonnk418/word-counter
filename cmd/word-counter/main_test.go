package main

import (
	"bytes"
	"testing"

	"github.com/eamonnk418/word-counter/internal/assert"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
		bytes    bool
		lines    bool
		words    bool
	}{
		{
			name:     "empty input",
			input:    "",
			expected: 0,
			bytes:    true,
			lines:    true,
			words:    true,
		},
		{
			name:     "bytes",
			input:    "hello",
			expected: 5,
			bytes:    true,
			lines:    false,
			words:    false,
		},
		{
			name:     "lines",
			input:    "hello\nworld",
			expected: 2,
			bytes:    false,
			lines:    true,
			words:    false,
		},
		{
			name:     "words",
			input:    "hello world",
			expected: 2,
			bytes:    false,
			lines:    false,
			words:    true,
		},
		{
			name:     "all",
			input:    "hello world",
			expected: 11,
			bytes:    true,
			lines:    true,
			words:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewBufferString(tc.input)

			got := count(r, tc.bytes, tc.lines, tc.words)

			assert.Equal(t, got, tc.expected)
		})
	}
}
