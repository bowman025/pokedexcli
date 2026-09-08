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
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "SOME   POKEMON   NAME",
			expected: []string{"some", "pokemon", "name"},
		},
		{
			input:    " Pikachu   ThEbest ",
			expected: []string{"pikachu", "thebest"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input %q: expected length %d, got %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("For input %q at index %d: expected %q, got %q", c.input, i, c.expected[i], actual[i])
			}

		}
	}
}
