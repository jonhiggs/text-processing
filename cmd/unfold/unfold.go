package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pborman/getopt/v2"
)

const (
	VERSION = "0.0.0"
)

//helpFlag := getopt.Bool('?', "display help")
//cmdFlag := getopt.StringLong("command", 'c', "default", "the command")

var (
	fileName = "/the/default/path"
	timeout  = time.Second * 5
	ver      bool
)

func init() {
	getopt.FlagLong(&ver, "version", 'V', "show version")
	getopt.FlagLong(&fileName, "path", 0, "the path")
	getopt.FlagLong(&timeout, "timeout", 't', "some timeout")
}

func main() {
	getopt.Parse()
	args := getopt.Args()

	fmt.Println(args)

	version()

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

func version() {
	fmt.Printf("unfold %s\n", VERSION)
	os.Exit(0)
}

func Usage() string {
	return fmt.Sprint("Whats this")
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
