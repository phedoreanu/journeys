package model

import "time"

type Driver struct {
	ID            string    `json:"id"`
	AvailableFrom time.Time `json:"available_from"`
	AvailableTill time.Time `json:"available_till"`
}
