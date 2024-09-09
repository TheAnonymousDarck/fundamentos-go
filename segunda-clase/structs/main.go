package structs

import "fmt"

type User struct {
	Name     string
	Email    string
	Address  string
	Birthday string
	phone    string
}

func main() {
	var davidUser User

	davidUser.Name = "David"
	davidUser.phone = "123456789"

	fmt.Println(davidUser)

	fmt.Println(davidUser.Name)
	fmt.Println(davidUser.phone)
}
