package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Alumno struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Course string `json:"course"`
	Grade  int    `json:"grade"`
	Id     int    `json:"id"`
}

type Data []Alumno

func main() {
	jsonFile, err := os.Open("segunda-clase/alumnos.json")
	if err != nil {
		fmt.Println("error al abrir el archivo", err)
		return
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("error al cerrar el archivo", err)
		}
	}(jsonFile)

	fmt.Println("el archivo se abrio exitosamente", jsonFile)

	bytesValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("error al leer el archivo", err)
		return
	}

	fmt.Println("el archivo se leyo exitosamente", string(bytesValue))

	var data Data
	err = json.Unmarshal(bytesValue, &data)
	if err != nil {
		fmt.Println("error al decodificar el archivo", err)
		return
	}

	//fmt.Println(data)

	promAge := 0
	for _, alumno := range data {
		fmt.Println(alumno.Name)
		fmt.Println(alumno.Age)
		fmt.Println(alumno.City)
		fmt.Println(alumno.Course)
		fmt.Println(alumno.Grade)
		fmt.Println(alumno.Id)
		fmt.Println("-------------------------")
		promAge += alumno.Age
	}

	fmt.Println("promedio de edades: ", promAge/len(data))
}
