package main

import "fmt"

func main() {
	for i := 12; i >= 0; i-- {
		switch i {
		case 10, 4, 2:
			continue
		default:
			fmt.Printf("i = %v\n", i)
		}
	}
}