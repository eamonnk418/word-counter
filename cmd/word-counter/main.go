package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	countBytes bool
	countLines bool
	countWords bool
)

func count(r io.Reader, countBytes, countLines, countWords bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countWords {
		scanner.Split(bufio.ScanWords)
	}

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func main() {
	flag.BoolVar(&countBytes, "b", false, "Count bytes instead of words")
	flag.BoolVar(&countLines, "l", false, "Count lines instead of words")
	flag.BoolVar(&countWords, "w", false, "Count words")
	flag.Parse()

	fmt.Println(count(os.Stdin, countBytes, countLines, countWords))
}
