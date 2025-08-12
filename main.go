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

		fmt.Println(line)

		tokens := make([]Token, 1)
		scanner := NewScanner(line)

		for {
			token := scanner.getNextToken()
			if token.Type == EOF {
				break
			}
			tokens = append(tokens, token)
		}

		fmt.Println("tokens -> ", tokens)
		for _, t := range tokens {
			fmt.Println(t.Type, t.Literal)
		}
	}
}
