package model

import "time"

type Journey struct {
	ID                string    `json:"id"`
	DepartureTime     time.Time `json:"departure_time"`
	DepartureLocation string    `json:"departure_location"`
	ArrivalTime       time.Time `json:"arrival_time"`
	ArrivalLocation   string    `json:"arrival_location"`
}
