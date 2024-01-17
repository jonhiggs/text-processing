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
		s = StripBlockquote(s)
		s = StripHeadings(s)
		s = StripCode(s)
		s = StripBold(s)
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// return true of string is a blockquote
func StripBlockquote(s string) string {
	return regexp.MustCompile(`^>\s*(.*)`).ReplaceAllString(s, `$1`)
}

// return true of string is a blockquote
func StripHeadings(s string) string {
	return regexp.MustCompile(`^#+\s*(.*)`).ReplaceAllString(s, `$1`)
}

// return the string without code
func StripCode(s string) string {
	return regexp.MustCompile("`[^`]*`").ReplaceAllString(s, `<CODE>`)
}

// return the string without code
func StripBold(s string) string {
	return regexp.MustCompile(`\*\*`).ReplaceAllString(s, "")
}
