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
				fmt.Printf("%s\n\n", p)
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

// take two strings and either unfold them into one string
func unfold(s, p string) string {
	return fmt.Sprintf("%s %s", p, strings.TrimSpace(s))
}

// return true if s should be folded into p
func isUnfoldable(s, p string) bool {
	if strings.TrimSpace(p) == "" {
		return false
	}

	if strings.HasPrefix(s, "-") {
		return false
	}

	return true
}
