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

	n := 0
	var buf string

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			flush(&buf, &n)
		} else {
			buf += fmt.Sprintf("%s\n", s)
		}
	}

	flush(&buf, &n)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func flush(s *string, n *int) {
	ss := sentences(*s)

	if len(ss) == 1 && strings.TrimSpace(ss[0]) == "" {
		return
	} else if len(ss) > 1 {
		*n++
		fmt.Printf("%2d  %s\n..  %s\n--\n", *n, ss[0], ss[len(ss)-1])
	} else {
		*n++
		fmt.Printf("%2d  %s\n--\n", *n, ss[0])
	}

	*s = ""
}

// from a paragraph of text, return the sentences.
func sentences(p string) []string {
	p = strings.TrimSpace(p)
	p = strings.Replace(p, ".\n", ". ", -1)
	str := strings.Split(p, ". ")

	for i, s := range str {
		if i != len(str)-1 {
			str[i] = fmt.Sprintf("%s.", s)
		}

		str[i] = strings.TrimSpace(str[i])
	}

	return str
}
