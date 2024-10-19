package usecases

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather?zipcode=01153000", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeatherHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	expectedKeys := []string{"temp_C", "temp_F", "temp_K"}
	for _, key := range expectedKeys {
		if !strings.Contains(body, key) {
			t.Errorf("handler returned unexpected body: missing key %v", key)
		}
	}

	var result map[string]float64
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Errorf("error parsing JSON response: %v", err)
	}

	for key, value := range result {
		truncatedValue := math.Round(value*10) / 10 // Arredonda para 1 casa decimal
		if value != truncatedValue {
			t.Errorf("value for %v has more than 1 decimal place: got %v", key, value)
		}
	}
}
