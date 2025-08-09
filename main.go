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
	}
}
