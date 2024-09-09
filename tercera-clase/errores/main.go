package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
			panic("Ups sistema parado")
		}
	}()
	fmt.Println("Sistema funcionando")
	fmt.Println("Programa parado")
}
