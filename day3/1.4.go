package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts,"123")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for linetext, filename := range counts {
		
			fmt.Printf("%s\t%s\n", linetext, filename)
		
	}
}

func countLines(f *os.File, counts map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "end"{
			break
		}
		counts[input.Text()]=arg
	}


	// NOTE: ignoring potential errors from input.Err()
}