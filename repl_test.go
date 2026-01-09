package main

import "testing"

func TestCleanInput(t *testing.T) {
	// ...
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "A Mild case of diarrhea BABY ",
			expected: []string{"a", "mild", "case", "of", "diarrhea", "baby"},
		},
		{
			input:    "whO tOld You to SCrub thEm",
			expected: []string{"who", "told", "you", "to", "scrub", "them"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected = %d; Actual = %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Expected = %s; Actual = %s", expectedWord, actual[i])
			}
		}
	}
}
