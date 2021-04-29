package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 3 {
		os.Exit(1)
	}

	left, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		os.Exit(1)
	}
	right, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		os.Exit(1)
	}

	result := rand.Int63n(right-left) + left
	fmt.Println(result)
}
