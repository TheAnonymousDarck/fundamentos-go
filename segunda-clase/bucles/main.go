package bucles

import "fmt"

func main() {
	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 2
	}

	//bucles sobre un arreglo
	numbers := []int{1, 2, 3, 4, 5}

	for i, numbers := range numbers {
		fmt.Println(i, numbers)
	}

	fmt.Println("tamaño del array:", len(numbers))

	for i := 0; i <= len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

}
