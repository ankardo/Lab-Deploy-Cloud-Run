package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/domain"
)

func TestGetLocationByCEP(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := "/ws/01153000/json/"
		if r.URL.String() != expectedURL {
			t.Fatalf("Expected URL to be %s, but got %s", expectedURL, r.URL.String())
		}

		location := domain.Location{
			ZipCode:      "01153-000",
			Street:       "Rua Teste",
			Complement:   "Apto 101",
			Neighborhood: "Bairro Exemplo",
			City:         "São Paulo",
			State:        "SP",
			IBGE:         "3550308",
			GIA:          "",
			DDD:          "11",
			SIAFI:        "7107",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(location)
	}))
	defer mockServer.Close()

	originalViaCEPUrl := viaCEPBaseURL
	viaCEPBaseURL = mockServer.URL
	defer func() { viaCEPBaseURL = originalViaCEPUrl }()

	location, err := GetLocationByCEP("01153000")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	expectedCity := "São Paulo"
	if location.City != expectedCity {
		t.Errorf("Expected city to be %s, but got %s", expectedCity, location.City)
	}

	expectedZipCode := "01153-000"
	if location.ZipCode != expectedZipCode {
		t.Errorf("Expected zipcode to be %s, but got %s", expectedZipCode, location.ZipCode)
	}
}
