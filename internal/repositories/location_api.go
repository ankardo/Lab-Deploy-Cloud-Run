package repositories

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/domain"
)

var viaCEPBaseURL = "https://viacep.com.br/ws/"

func GetLocationByCEP(zipcode string) (*domain.Location, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location domain.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.ZipCode == "" {
		return nil, errors.New("invalid zipcode")
	}

	return &location, nil
}
