package monitor

import (
	"encoding/json"
	"fmt"
	"io"
	config "meliprices/config"
	"meliprices/database"
	"meliprices/services"
	"net/http"
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
	response, err := p.httpClient.Get(config.MELI_URL + "/items/" + p.Id)
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

func (p *PriceMonitor) Save(service services.PriceService) error {
	price := database.Price{
		IdProduct:     p.Id,
		Price:         p.Price,
		BasePrice:     p.BasePrice,
		OriginalPrice: p.OriginalPrice,
		DateTime:      p.DateTime,
	}

	err := service.CreatePrice(price)

	if err != nil {
		fmt.Println("Error al crear el precio", err)
	}

	fmt.Println("Se guardo el precio")

	return nil
}

func NewPriceMonitor(id string, client *http.Client) *PriceMonitor {
	return &PriceMonitor{Id: id, httpClient: http.DefaultClient}
}
