package main

import "fmt"

func main() {

	c := make(chan int, 1)
	c2 := make(chan bool, 1)
	value := 0
	go func() {
		value = 100
		c <- 42
	}()
	go func() {
		value = 200
		c2 <- true
	}()

	//value := <-c
	<-c2
	fmt.Println(value)
}
