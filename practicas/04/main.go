package main

import (
	"encoding/json"
	"fmt"
	"io"
	"meliitems/database"
	"meliitems/services"
	"net/http"
	"net/url"
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
	db, err := database.NewDatabaseDriver()

	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return
	}

	fmt.Println("Conexión exitosa")
	err = db.AutoMigrate(&database.Item{})
	if err != nil {
		return
	}
	itemService := services.NewItemService(db)

	client := &http.Client{}
	q := "Motorola"
	resp, err := client.Get("https://api.mercadolibre.com/sites/MLM/search?q=" + url.QueryEscape(q))

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error", err)
		}
	}(resp.Body)
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
		item := database.Item{
			Id:         item.Id,
			Title:      item.Title,
			Price:      item.Price,
			Condition:  item.Condition,
			CategoryId: item.CategoryId,
			Url:        item.Url,
		}

		err = itemService.CreateItem(item)
		if err != nil {
			fmt.Println("Error al crear el item", err)
		}
	}

	fmt.Println("Items creados exitosamente en la base de datos")
}
