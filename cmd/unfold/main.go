package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	VERSION = "v0.1.0"
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
				fmt.Println(fmt.Errorf("unfold: %s: %s", a, errors.Unwrap(err)))
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
		processFile(f)
	}
}

func processFile(f *os.File) {
	scanner := bufio.NewScanner(f)
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
	if err := scanner.Err(); err != nil {
		fmt.Println(fmt.Errorf("unfold: %s", err))
		os.Exit(1)
	}

	fmt.Printf("%s\n", p)
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

func printHelp() {
	h := `Usage: unfold [FILE]...
Combine lines, reversing the affects of fold(1).

With no FILE, read standard input.

  --version    print the version and exit
  --help       display help and exit
`
	fmt.Println(h)
}

func printVersion() {
	fmt.Printf("unfold %s\n", VERSION)
}
