package main

import (
	"fmt"
	"time"
)

func printWord(word string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func main() {
	go printWord("hello", 5)
	printWord("world", 5)
	time.Sleep(1 * time.Second)
}
