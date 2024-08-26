package main

import "fmt"

func main() {
	name := "david" + "Resendiz"
	var age int = 23
	const pi float32 = 3.1415

	fmt.Printf("Name: %v, Age: %v, Pi: %v, IsPublic: %v", name, age, pi, false)

	total := 10 + 1
	fmt.Println("Total: ", total)

	total += 10
	fmt.Println("suma: ", total)

	total -= 10
	fmt.Println("resta: ", total)

	total /= 10
	fmt.Println("division: ", total)

	total *= 10
	fmt.Println("multiplicacion: ", total)

	total %= 100
	fmt.Println("modulo: ", total)

	total++
	fmt.Println(total)

	result := 10 != 11
	fmt.Println(result)

	// if normal
	if result {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	// if and
	if 10 > 5 && 10 < 20 {
		fmt.Println("10 es mayor a 5")
	} else {
		fmt.Println("10 no es mayor a 5")
	}

	// if or
	if 10 > 5 || 10 < 20 {
		fmt.Println("10 es mayor a 5")
	} else {
		fmt.Println("10 no es mayor a 5")
	}

	// switch
	day := get_day_name(3)

	switch day {
	case "lunes":
		fmt.Println("Hoy es lunes")
	case "martes":
		fmt.Println("Hoy es martes")
	default:
		fmt.Println("Hoy no es lunes ni martes")
	}

	if day := get_day_name(1); day == "lunes" {
		fmt.Println("Hoy es lunes")
	}

}

func get_day_name(day_number int) string {
	switch day_number {
	case 1:
		return "Lunes"
	case 2:
		return "Martes"
	case 3:
		return "Miercoles"
	case 4:
		return "Jueves"
	case 5:
		return "Viernes"
	case 6:
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Dia no valido"
	}
}
