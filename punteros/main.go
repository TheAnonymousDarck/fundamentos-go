package main

import "fmt"

func changeName(name *string) {
	*name = "Juan"
}

func main() {
	name := "David"
	ptr := &name
	fmt.Println(ptr)
	fmt.Println(*ptr)

	changeName(&name)
	fmt.Println(name)
}
