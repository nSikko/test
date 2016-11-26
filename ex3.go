package main

import (
    "fmt"
    "math/rand"
    // "strconv"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Test", rand.Intn(999)) // 0 < i <= 999
}

func main() {
	init()
}