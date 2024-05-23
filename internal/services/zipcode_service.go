package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rzeradev/google-cloud-run/configs"
)

type Location struct {
	City string `json:"localidade"`
}

func FetchLocation(zipcode string) (*Location, error) {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(cfg.CepAPIURL, zipcode)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("invalid response from viacep")
	}
	defer resp.Body.Close()

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, errors.New("failed to decode location response")
	}

	if location.City == "" {
		return nil, errors.New("city not found for the given zipcode")
	}

	return &location, nil
}
