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

func TestStripBoldAndItalic(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"this is *italic*", "this is italic"},
		{"this is **bold**", "this is bold"},
		{"this is ***bold and italic***", "this is bold and italic"},
		{"* this is a bullet", "* this is a bullet"},
	}

	for _, test := range tests {
		got := stripBoldAndItalic(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripLink(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"[x](y)", "x"},
		{"one [two](three)", "one two"},
	}

	for _, test := range tests {
		got := stripLink(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripImage(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"![x](y)", "[x]"},
		{"one ![two](three)", "one [two]"},
		{"[x](y)", "[x](y)"},
	}

	for _, test := range tests {
		got := stripImage(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}
