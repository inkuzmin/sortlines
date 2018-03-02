package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"sort"
	"flag"
)

func main() {

	input := flag.String("i", "", "Input file in a tab-separated format (default: stdin)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	if *input != "" {
		inFile, _ := os.Open(*input)
		defer inFile.Close()

		scanner = bufio.NewScanner(inFile)
	}

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		sort.Strings(words)
		for i, word := range words {
			fmt.Print(word)
			if i != len(words) - 1 {
				fmt.Print("\t")
			} else {
				fmt.Print("\n")
			}
		}
	}

}
