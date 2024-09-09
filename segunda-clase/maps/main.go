package maps

import "fmt"

func main() {
	alumnos := map[int]string{
		290391: "David",
		290392: "Carlos",
		290393: "Andres",
		290394: "Andrea",
		290395: "Isabel",
	}

	fmt.Println(alumnos)
	alumnos[290395] = "Isabel agregado"
	fmt.Println(alumnos[290395])

	delete(alumnos, 290395)
	fmt.Println(alumnos)

	//	saber si existe una llave
	nombre, existe := alumnos[290395]
	fmt.Println(nombre, existe)

	_, ok := alumnos[290395]
	if ok {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
}
