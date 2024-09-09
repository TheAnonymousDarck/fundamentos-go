package arreglos

import "fmt"

func main() {
	var numbers [5]int
	numbers2 := [5]int{1, 2, 3, 4, 5}
	numbers[0] = 10

	fmt.Println(numbers)
	fmt.Println(numbers2)

	//	areglo dinamico
	var dynNumbers []int
	dynNumbers = append(dynNumbers, 1)
	dynNumbers = append(dynNumbers, 2)
	dynNumbers = append(dynNumbers, 3)
	fmt.Println(dynNumbers)

	//	areglos multidimensionales Matrices
	matrix := [2][3]int{{1, 2}, {4, 5, 6}}
	fmt.Println(matrix[0][1])

	//	comparacion de arreglos
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	if a == b {
		fmt.Println("Iguales")
	}

	//	slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice[1:3])
	fmt.Println(slice[:3])
	fmt.Println(slice[3:])

}
