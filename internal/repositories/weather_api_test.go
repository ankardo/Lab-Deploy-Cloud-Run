package repositories

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ankardo/Lab-Deploy-Cloud-Run/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetWeatherByLocation(t *testing.T) {
	mockWeatherData := `{
		"current": {
			"temp_c": 18.25
		}
	}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockWeatherData))
	}))
	defer mockServer.Close()

	client := mockServer.Client()

	location := &domain.Location{
		City: "São Paulo",
	}

	weather, err := GetWeatherByLocation(location, client)
	assert.Nil(t, err)
	assert.NotNil(t, weather)
	assert.Equal(t, 18.3, weather.TempC)
	assert.Equal(t, 64.9, weather.TempF)
	assert.Equal(t, 291.4, weather.TempK)
}
