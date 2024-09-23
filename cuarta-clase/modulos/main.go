package modulos

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"modulos/config"
	"modulos/monitor"
	"net/http"
)

func main() {
	cronJob := cron.New()
	client := &http.Client{}

	_, err := cronJob.AddFunc(config.TIMER, func() {
		println("Se ejecuta cada 2 segundos")
		priceMonitor := monitor.NewPriceMonitor("MLM2590077782", client)

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
