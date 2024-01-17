package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestIsUnfoldable(t *testing.T) {

	var tests = []struct {
		p    string
		s    string
		want bool
	}{
		{"", "A Heading", false},
		{"one", "two", true},
		{"- item", "- item", false},
		{"hyphenated-", "word", true},
		{"", "", false},
	}

	// The execution loop
	for _, test := range tests {
		got := isUnfoldable(test.s, test.p)
		if got != test.want {
			t.Errorf("'%s' and '%s' = %v", test.s, test.p, test.want)
		}
	}

}
