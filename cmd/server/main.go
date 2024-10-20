package main

import (
	"net/http"
	"os"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/configs"
	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/usecases"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.WebServerPort
	}
	http.HandleFunc("/weather", usecases.GetWeatherHandler)
	http.ListenAndServe(":"+port, nil)
}
