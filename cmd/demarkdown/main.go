package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		s = stripBlockquote(s)
		s = stripHeading(s)
		s = stripHorizontalRule(s)
		s = stripCode(s)
		s = stripBoldAndItalic(s)
		s = stripImage(s)
		s = stripLink(s)
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
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
	return regexp.MustCompile("`[^`]*`").ReplaceAllString(s, `<CODE>`)
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
	return regexp.MustCompile(`!\[([^]]*)\]\(.*\)`).ReplaceAllString(s, `[$1]`)
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
