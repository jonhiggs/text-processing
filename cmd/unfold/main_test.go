package main

import (
	"testing"
)

func TestIsUnfoldable(t *testing.T) {
	var tests = []struct {
		p    string
		s    string
		want bool
	}{
		{"", "A Heading", false},
		{"one", "two", true},
		{"- item", "- item", false},
		{"word", "", false},
		{"hyphenated-", "word", true},
		{"", "", false},
	}

	for _, test := range tests {
		got := isUnfoldable(test.s, test.p)
		if got != test.want {
			t.Errorf("'%s' and '%s' = %v", test.s, test.p, test.want)
		}
	}

}

func TestUnfold(t *testing.T) {
	var tests = []struct {
		p    string
		s    string
		want string
	}{
		{"one", "two", "one two"},
		{"   one", "  two    ", "   one two"},
	}

	for _, test := range tests {
		got := unfold(test.s, test.p)
		if got != test.want {
			t.Errorf("'%s' and '%s' = '%s', got '%s'", test.s, test.p, test.want, got)
		}
	}

}
