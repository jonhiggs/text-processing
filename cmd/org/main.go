package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	VERSION = "0.0.0"
)

var (
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
		case "-v", "--version":
			optVer = true
		case "-h", "--help":
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
				fmt.Println(fmt.Errorf("org: %s: %s", a, errors.Unwrap(err)))
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

func printHelp() {
	h := `Usage: org [FILE]...
Print the first and last sentence from each paragraph.

With no FILE, read standard input.

  --version    print the version and exit
  --help       display help and exit
`
	fmt.Println(h)
}

func printVersion() {
	fmt.Printf("unfold %s\n", VERSION)
}
