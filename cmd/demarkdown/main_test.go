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

// must strip code fences before stripping code
func TestStripCode(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"`code`", "<....>"},
		{"one `code` and two `codes`.", "one <....> and two <.....>."},
		{"a string with `code`", "a string with <....>"},
		//{"```python", "```python"},
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
		{"![](y)", "[img]"},
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

func TestStripEscapeSequences(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"an asterisk (\\*)", "an asterisk (*)"},
		{"a dash (\\-)", "a dash (-)"},
		{"a backslash (\\\\)", "a backslash (\\)"},
		{"a backtick (\\`)", "a backtick (`)"},
		{"an underscore (\\_)", "an underscore (_)"},
		{"a curly left brace (\\{)", "a curly left brace ({)"},
		{"a curly right brace (\\})", "a curly right brace (})"},
		{"a left bracket (\\[)", "a left bracket ([)"},
		{"a right bracket (\\])", "a right bracket (])"},
		{"a left angle bracket (\\<)", "a left angle bracket (<)"},
		{"a right angle bracket (\\>)", "a right angle bracket (>)"},
		{"a left parentheses (\\()", "a left parentheses (()"},
		{"a right parentheses (\\))", "a right parentheses ())"},
		{"a hash (\\#)", "a hash (#)"},
		{"a plus (\\+)", "a plus (+)"},
		{"a minus (\\-)", "a minus (-)"},
		{"a dot (\\.)", "a dot (.)"},
		{"an exclamation (\\!)", "an exclamation (!)"},
		{"a pipe (\\|)", "a pipe (|)"},
	}

	for _, test := range tests {
		got := stripEscapeSequences(test.s)
		if got != test.want {
			t.Errorf("'%s' = '%s', got '%s'", test.s, test.want, got)
		}
	}
}

func TestIsFence(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"```", true},
		{"```shell", true},
		{"  ```go", true},
	}

	for _, test := range tests {
		got := isFence(test.s)
		if got != test.want {
			t.Errorf("'%s' = %v", test.s, test.want)
		}
	}
}
