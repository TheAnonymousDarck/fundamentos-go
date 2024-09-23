package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"meliprices/config"
	"meliprices/database"
	"meliprices/monitor"
	"meliprices/services"
	"net/http"
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
	err = db.AutoMigrate(&database.Price{})
	priceService := services.NewPriceService(db)

	if err != nil {
		return
	}

	client := &http.Client{}
	cronJob := cron.New()
	client = &http.Client{}

	_, err = cronJob.AddFunc(config.TIMER, func() {
		println("Se ejecuta cada 2 segundos")
		priceMonitor := monitor.NewPriceMonitor("MLM2590077782", client)

		err := priceMonitor.Fetch()
		if err != nil {
			return
		}

		err = priceMonitor.Save(priceService)
		if err != nil {
			return
		}

		println("Precio actual: ", priceMonitor.Price)

		err = priceMonitor.Fetch()
		if err != nil {
			fmt.Println("error al obtener el precio del item")
		}
	})
	if err != nil {
		return
	}
	cronJob.Start()

	select {}
}
