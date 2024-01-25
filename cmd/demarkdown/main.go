package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	VERSION = "v0.1.0"
)

var (
	fenced  = false
	files   []*os.File
	optHelp = false
	optVer  = false
)

func init() {
	for i, a := range os.Args {
		if i == 0 {
			continue
		}

		// don't trip up when running the tests
		if strings.HasPrefix(a, "-test.") {
			continue
		}

		switch a {
		case "--version":
			optVer = true
		case "--help":
			optHelp = true
		case "--":
		default:
			if strings.HasPrefix(a, "-") {
				fmt.Printf("Unsupported flag '%s'.\n\n", a)
				printHelp()
				os.Exit(1)
			}

			f, err := os.Open(a)
			if err != nil {
				fmt.Println(fmt.Errorf("demarkdown: %s: %s", a, errors.Unwrap(err)))
				os.Exit(1)
			}

			files = append(files, f)
		}
	}

	if optHelp {
		printHelp()
		os.Exit(0)
	}

	if optVer {
		printVersion()
		os.Exit(0)
	}

	if len(files) == 0 {
		files = append(files, os.Stdin)
	}
}

func main() {
	for _, f := range files {
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			s := scanner.Text()

			if isFenced(s) {
				s = ""
			} else {
				s = stripBlockquote(s)
				s = stripHeading(s)
				s = stripHorizontalRule(s)
				s = stripCodeFence(s)
				s = stripCode(s)
				s = stripBoldAndItalic(s)
				s = stripImage(s)
				s = stripLink(s)
				s = stripLinkReference(s)
				s = stripEscapeSequences(s)
			}

			fmt.Println(s)
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
}

// return true of string is a blockquote
func stripBlockquote(s string) string {
	return regexp.MustCompile(`^>\s*(.*)`).ReplaceAllString(s, `$1`)
}

// return true of string is a blockquote
func stripHeading(s string) string {
	return regexp.MustCompile(`^#+\s*(.*)`).ReplaceAllString(s, `$1`)
}

// return the string without code
func stripCode(s string) string {
	inCode := false
	var runes []rune

	for _, char := range s {
		if char == '`' {
			if inCode {
				runes = append(runes, '>')
			} else {
				runes = append(runes, '<')
			}

			inCode = !inCode
		} else {
			if inCode {
				runes = append(runes, '.')
			} else {
				runes = append(runes, char)
			}
		}
	}

	return string(runes)
}

// return string without code fence
func stripCodeFence(s string) string {
	if isFence(s) {
		return ""
	} else {
		return s
	}
}

// return the string without bold or italic tags
func stripBoldAndItalic(s string) string {
	s = regexp.MustCompile(`\*{3}([^\*]+)\*{3}`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`\*{2}([^\*]+)\*{2}`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`\*([^\*]+)\*`).ReplaceAllString(s, `$1`)

	s = regexp.MustCompile(`_{3}([^_]+)_{3}`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`_{2}([^_]+)_{2}`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`_([^_]+)_`).ReplaceAllString(s, `$1`)
	return s
}

// return the string without links
func stripLink(s string) string {
	// make sure it's not an image link
	s = regexp.MustCompile(`([^!])\[([^]]*)\]\([^)]*\)`).ReplaceAllString(s, `$1$2`)
	s = regexp.MustCompile(`([^!])\[([^]]*)\]\ ?\[[^]]*\]`).ReplaceAllString(s, `$1$2`)

	// but also handle links at the start of the line
	s = regexp.MustCompile(`^\[([^]]*)\]\([^)]*\)`).ReplaceAllString(s, `$1`)
	s = regexp.MustCompile(`^\[([^]]*)\]\ ?\[[^]]*\]`).ReplaceAllString(s, `$1`)
	return s
}

// return the string without link references
func stripLinkReference(s string) string {
	if regexp.MustCompile(`^\[[^]]+\]:\ `).MatchString(s) {
		return ""
	}

	return s
}

// return the string without images
func stripImage(s string) string {
	s = regexp.MustCompile(`!\[([^]]+)\]\(.*\)`).ReplaceAllString(s, `[$1]`)
	s = regexp.MustCompile(`!\[\]\(.*\)`).ReplaceAllString(s, `[img]`)

	return s
}

// return the string without horizontal lines.
// This also removes the underline from headings when in the format of:
//
//	Heading
//	-------
func stripHorizontalRule(s string) string {
	if regexp.MustCompile(`^={3,}$`).MatchString(s) {
		return ""
	}

	if regexp.MustCompile(`^-{3,}$`).MatchString(s) {
		return ""
	}

	if regexp.MustCompile(`^\*{3,}$`).MatchString(s) {
		return ""
	}

	return s
}

// return string without the escape sequences
func stripEscapeSequences(s string) string {
	s = strings.ReplaceAll(s, `\*`, `*`)
	s = strings.ReplaceAll(s, `\-`, `-`)
	s = strings.ReplaceAll(s, "\\`", "`")
	s = strings.ReplaceAll(s, `\_`, `_`)
	s = strings.ReplaceAll(s, `\{`, `{`)
	s = strings.ReplaceAll(s, `\}`, `}`)
	s = strings.ReplaceAll(s, `\[`, `[`)
	s = strings.ReplaceAll(s, `\]`, `]`)
	s = strings.ReplaceAll(s, `\<`, `<`)
	s = strings.ReplaceAll(s, `\>`, `>`)
	s = strings.ReplaceAll(s, `\(`, `(`)
	s = strings.ReplaceAll(s, `\)`, `)`)
	s = strings.ReplaceAll(s, `\#`, `#`)
	s = strings.ReplaceAll(s, `\+`, `+`)
	s = strings.ReplaceAll(s, `\-`, `-`)
	s = strings.ReplaceAll(s, `\.`, `.`)
	s = strings.ReplaceAll(s, `\!`, `!`)
	s = strings.ReplaceAll(s, `\|`, `|`)
	s = strings.ReplaceAll(s, `\\`, `\`)
	return s
}

// tracks the state of whether the current line is within a code fence.
func isFenced(s string) bool {
	if isFence(s) {
		fenced = !fenced
	}

	return fenced
}

// checks if the supplied string is detected as a fence.
func isFence(s string) bool {
	return regexp.MustCompile("^`{3}").MatchString(strings.TrimSpace(s))
}

func printHelp() {
	h := `Usage: demarkdown [FILE]...
Remove formatting from Markdown documents.

With no FILE, read standard input.

      --help            display help and exit
      --version         print the version and exit
`
	fmt.Println(h)
}

func printVersion() {
	fmt.Printf("demarkdown %s\n", VERSION)
}
