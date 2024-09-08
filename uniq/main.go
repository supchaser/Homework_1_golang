package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}

		prev = txt
		fmt.Fprintln(output, txt)
	}
}

func main() {
	uniq(os.Stdin, os.Stdout)
}
