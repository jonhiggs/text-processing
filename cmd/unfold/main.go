package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var p string

	n := 0
	for scanner.Scan() {
		s := scanner.Text()
		if n == 0 {
			p = s
		} else {
			if isUnfoldable(p, s) {
				p = unfold(s, p)
			} else {
				fmt.Printf("%s\n", p)
				p = s
			}
		}

		n++
	}

	fmt.Printf("%s\n", p)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// take two strings and unfold them into one
func unfold(s, p string) string {
	return fmt.Sprintf("%s %s", p, strings.TrimSpace(s))
}

// return true if s should be unfolded into p
func isUnfoldable(s, p string) bool {
	if strings.TrimSpace(p) == "" && strings.TrimSpace(s) == "" {
		return true
	}

	if strings.TrimSpace(p) == "" {
		return false
	}

	if strings.TrimSpace(s) == "" {
		return false
	}

	if strings.HasPrefix(s, "-") {
		return false
	}

	return true
}
