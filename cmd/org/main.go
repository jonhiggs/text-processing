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

	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// from a paragraph of text, return the sentences.
func sentences(p string) []string {
	str := strings.Split(p, ". ")

	for i, s := range str {
		str[i] = fmt.Sprintf("%s.", s)
	}

	return str
}
