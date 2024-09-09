package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"net/http"
	"os"
	"time"
)

type PriceMonitor struct {
	Id            string
	Price         float64 `json:"price"`
	BasePrice     float64 `json:"base_price"`
	OriginalPrice float64 `json:"original_price"`
	DateTime      string  `json:"date_time"`
	httpClient    *http.Client
}

func (p *PriceMonitor) Fetch() error {
	response, err := p.httpClient.Get("https://api.mercadolibre.com/items/" + p.Id)
	if err != nil {
		fmt.Println("error al obtener el precio del item")
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error al cerrar el body")
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error al leer el body")
		return err
	}

	err = json.Unmarshal(body, &p)
	p.DateTime = time.Now().UTC().String()
	if err != nil {
		return err
	}

	return nil
}

func (p *PriceMonitor) Save() error {
	jsonFile, erro := os.OpenFile("prices.jsonl",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if erro != nil {
		fmt.Println("error al abrir el json")
		return erro
	}
	defer jsonFile.Close()

	jsonData, err := json.Marshal(p)
	if err != nil {
		return err
	}

	jsonFile.Write(jsonData)
	jsonFile.WriteString("\n")

	fmt.Println("Se guardo el precio")

	return nil
}

func main() {
	client := &http.Client{}
	cronJob := cron.New()

	_, err := cronJob.AddFunc("@every 2s", func() {
		println("Se ejecuta cada 2 segundos")
		priceMonitor := PriceMonitor{
			Id:         "MLM2590077782",
			httpClient: client,
		}

		err := priceMonitor.Fetch()
		if err != nil {
			return
		}

		err = priceMonitor.Save()
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
