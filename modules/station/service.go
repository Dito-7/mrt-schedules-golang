package station

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Dito-7/mrt-schedules-golang/common/client"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw API Response: %s\n", string(byteResponse))

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	for _, station := range stations {
		response = append(response, StationResponse{
			NID:   station.NID,
			Title: station.Title,
		})
	}

	return
}
