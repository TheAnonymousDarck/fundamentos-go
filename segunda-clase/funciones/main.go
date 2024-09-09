package funciones

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	defer func() {
		fmt.Println("se ejecuta al final de main")
		fmt.Println("esto se ejecuta tambien al final")
	}()

	nameComplete := test("David", 23)
	fmt.Println(nameComplete)

	nameComplete2, err := test2("David", 23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nameComplete2)

	day, err := get_day_name(5)
	if err != nil {
		fmt.Println("", err)
	}
	fmt.Println(day)

	fmt.Println(suma("", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//	función anonima
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(9, 10))

	//	funcion anonima autoejecutable
	func() {
		fmt.Println("Esta es una funcion anonima que se ejecuta automaticamente")
	}()

}

func test(name string, age int) string {
	return fmt.Sprintf("Name: %v, Age: %v", name, age)
}

// funciones con retorno multiple
func test2(name string, age int) (string, error) {
	return name + " Edad: " + strconv.Itoa(age), nil
}

func get_day_name(day_number int) (string, error) {
	switch day_number {
	case 1:
		return "Lunes", nil
	case 2:
		return "Martes", nil
	case 3:
		return "Miercoles", nil
	case 4:
		return "Jueves", nil
	case 5:
		return "Viernes", nil
	case 6:
		return "Sabado", nil
	case 7:
		return "Domingo", nil
	default:
		return "", errors.New("dia no valido")
	}
}

// funciones con parametros multiples
func suma(text string, numbers ...int) int {
	total := 0
	print(text)
	for _, number := range numbers {
		total += number
	}
	return total
}
