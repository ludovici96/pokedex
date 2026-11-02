package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello   ",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "one two three",
			expected: []string{"one", "two", "three"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"Expected %d words, got %d for input %q",
				len(c.expected),
				len(actual),
				c.input,
			)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(
					"Expected %q, got %q for input %q",
					c.expected[i],
					actual[i],
					c.input,
				)
			}
		}
	}
}
