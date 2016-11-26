package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Gate", "Comet", "Pizza"}
	fmt.Println(words)

	for _, word := range words {
		if strings.Contains(word, "e") {
			fmt.Printf("Found one! %s\n", string(word[0]) )
		}
	}
		
}