package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("mini-db >")

		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.EqualFold(line, "exit") {
			break
		}

		tokens := make([]Token, 0)
		scanner := NewScanner(line)

		longestLen := 0

		for {
			token := scanner.getNextToken()
			if token.Type == EOF {
				break
			}

			if len(token.Type) > longestLen {
				longestLen = len(token.Type)
			}

			tokens = append(tokens, token)
		}

		for _, t := range tokens {
			spaces := longestLen - len(t.Type)
			fmt.Printf("type: %v %*s,  literal: %v \n", t.Type, spaces, "", t.Literal)
		}
	}
}
