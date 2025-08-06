package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func printHelp() {
	fmt.Println(`--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.`)
}

func main() {
	var input string
	insert := ""
	shouldOrder := false

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	for _, arg := range os.Args[1:] {
		switch {
		case arg == "--help" || arg == "-h":
			printHelp()
			return
		case strings.HasPrefix(arg, "--insert="):
			insert = strings.TrimPrefix(arg, "--insert=")
		case strings.HasPrefix(arg, "-i="):
			insert = strings.TrimPrefix(arg, "-i=")
		case arg == "--order" || arg == "-o":
			shouldOrder = true
		default:
			input += arg
		}
	}

	// Append insert string if present
	input += insert

	// Sort if --order is set
	if shouldOrder {
		chars := strings.Split(input, "")
		sort.Strings(chars)
		input = strings.Join(chars, "")
	}

	fmt.Println(input)
}
