package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rzeradev/google-cloud-run/configs"
	"github.com/rzeradev/google-cloud-run/internal/models"
	"github.com/rzeradev/google-cloud-run/pkg/utils"
)

func FetchWeather(city string) (*models.Weather, error) {
	url := fmt.Sprintf(configs.Cfg.WeatherAPIURL, configs.Cfg.WeatherAPIKey, city)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("invalid response from weatherapi")
	}
	defer resp.Body.Close()

	var apiResponse struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, errors.New("failed to decode weather response")
	}

	tempC := apiResponse.Current.TempC
	weather := &models.Weather{
		TempC: tempC,
		TempF: utils.CelsiusToFahrenheit(tempC),
		TempK: utils.CelsiusToKelvin(tempC),
	}

	return weather, nil
}
