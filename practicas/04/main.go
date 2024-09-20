package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

type Item struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Condition  string  `json:"condition"`
	CategoryId string  `json:"category_id"`
	Url        string  `json:"url"`
}

type Response struct {
	Items []Item `json:"results"`
}

func main() {
	db,err := databse.newDatabaseDriver()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return nil, err
	}
	fmt.Println("Conexión exitosa")
	db.AutoMigrate(&database.Item{})
	itemService := services.NewItemService(db)
	err = itemService.CreateItem(database.Item{
		Id: "MLM123456",
		Title: "Motorola G6",
		Price: 5000.00,
	})

	if err != nil {
		fmt.Println("Error al crear el item: ", err)
	}
	return


	q := "Motorola"
	resp, err := client.Get("https://api.mercadolibre.com/sites/MLM/search?q=" + url.QueryEscape(q))

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Response status:", resp.StatusCode)
	fmt.Println("Body:", resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	//fmt.Println("Body:", string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	for _, item := range response.Items {
		fmt.Println("id:", item.Id)
	}

	//	Creamos el JSON
	file, err := os.Create("items.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fileData, err := json.Marshal(response.Items)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	_, err = file.Write(fileData)

	if err != nil {
		fmt.Println("Error al escribir el archivo", err)
		return
	}

}
