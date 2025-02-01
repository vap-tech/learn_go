package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := input.Text()
		if data == "exit" {
			break
		}
		counts[data]++
	}
}
