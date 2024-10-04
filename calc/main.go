package main

import (
	"bufio"
	"calc/shuntingYardAlg"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, I am a calculator!")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tokens := shuntingYardAlg.Tokenize(scanner.Text())
		result, err := shuntingYardAlg.ShuntingYardAlg(tokens)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
