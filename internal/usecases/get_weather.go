package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/domain"
	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/repositories"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("zipcode")
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := repositories.GetLocationByCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := repositories.GetWeatherByLocation(location, http.DefaultClient)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("could not get information about weather, error: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.NewWeather(weather.TempC))
}
