package main

import (
	"testing"
)

func TestStripBlockquote(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"nothing", "nothing"},
		{"> quote", "quote"},
		{">quote", "quote"},
		{">    quote", "quote"},
		{"  >    quote", "  >    quote"},
	}

	for _, test := range tests {
		got := stripBlockquote(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripHeading(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"# h1", "h1"},
		{"## h2", "h2"},
		{"### h3", "h3"},
		{"not heading", "not heading"},
	}

	for _, test := range tests {
		got := stripHeading(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripCode(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"`code`", "<CODE>"},
		{"a string with `code`", "a string with <CODE>"},
	}

	for _, test := range tests {
		got := stripCode(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripBold(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"this is **bold**", "this is bold"},
	}

	for _, test := range tests {
		got := stripBold(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}
