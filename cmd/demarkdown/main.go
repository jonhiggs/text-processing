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
		s = stripCode(s)
		s = stripBold(s)
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

// return the string without bold tags
func stripBold(s string) string {
	return regexp.MustCompile(`\*\*`).ReplaceAllString(s, "")
}
