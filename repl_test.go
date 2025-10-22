package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "Hello, \nWorld!",
			expected: []string{"hello,", "world!"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of slice does not match: expected %d/ actual %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %v does not match: expected %s/ actual %s", i, expectedWord, word)
			}
		}
	}
}
