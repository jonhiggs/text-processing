package main

import (
	"testing"
)

func TestSentences(t *testing.T) {
	var tests = []struct {
		s         string
		wantFirst string
		wantLast  string
		wantLen   int
	}{
		{"no dot", "no dot", "no dot", 1},
		{"one. two", "one.", "two", 2},
		{"one. two.", "one.", "two.", 2},
		{"one.  two.", "one.", "two.", 2},
		{"one. two. three.", "one.", "three.", 3},
		{"one.\ntwo.\nthree.", "one.", "three.", 3},
		{"one\ntwo", "one two", "one two", 1},
		{"one.\ntwo.\nthree.\n\n\n\n", "one.", "three.", 3},
		{"-a\n-b\n-c\n", "-a -b -c", "-a -b -c", 1},
		{"* one\n* two\n", "* one * two", "* one * two", 1},
		{"* one.\n* two.\n", "* one.", "* two.", 2},
	}

	for _, test := range tests {
		got := sentences(test.s)

		if got[0] != test.wantFirst {
			t.Errorf("first = '%s', got '%s'", test.wantFirst, got[0])
		}

		if got[len(got)-1] != test.wantLast {
			t.Errorf("last = '%s', got '%s'", test.wantLast, got[len(got)-1])
		}

		if len(got) != test.wantLen {
			t.Errorf("length = %d, got %d", test.wantLen, len(got))
		}
	}
}
