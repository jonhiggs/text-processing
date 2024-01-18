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
		{"```python", "```python"},
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
		{"this is _italic_", "this is italic"},
		{"this is __bold__", "this is bold"},
		{"this is ___bold and italic___", "this is bold and italic"},
		{"one __bold__ and one _italic_", "one bold and one italic"},
		{"this **is bold** and so is **this**", "this is bold and so is this"},
		{"_ this is a bullet", "_ this is a bullet"},
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
		{"[section](#section_link)", "section"},
		{"a [link one](1), and a [link two](2)", "a link one, and a link two"},
		{"a [reference link][a]", "a reference link"},
		{"a [reference link] [a] with a space", "a reference link with a space"},
		{"[reference link a][a] and [reference link b] [b]", "reference link a and reference link b"},
	}

	for _, test := range tests {
		got := stripLink(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestStripLinkReference(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"[1]: https://abc.com#Lifestyle", ""},
		{"[1]: https://abc.com#Lifestyle \"blah blah\"", ""},
		{"[1]: https://abc.com#Lifestyle 'blah blah'", ""},
		{"[1]: https://abc.com#Lifestyle (blah blah)", ""},
		{"[1]: <https://abc.com/blah/Blah#Lifestyle> \"blah blah\"", ""},
		{"[1]: <https://abc.com/blah/Blah#Lifestyle> 'blah blah'", ""},
		{"[1]: <https://abc.com/blah/Blah#Lifestyle> (blah blah)", ""},
	}

	for _, test := range tests {
		got := stripLinkReference(test.s)
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

func TestHorizontalRule(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"---", ""},
		{"===", ""},
		{"***", ""},
		{"--", "--"},
		{"==", "=="},
		{"**", "**"},
	}

	for _, test := range tests {
		got := stripHorizontalRule(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}
