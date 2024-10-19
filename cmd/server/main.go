package main

import (
	"net/http"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/configs"
	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/usecases"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/weather", usecases.GetWeatherHandler)
	http.ListenAndServe(":"+cfg.WebServerPort, nil)
}
