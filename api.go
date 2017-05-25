package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/phedoreanu/journeys/model"
)

func FindByDriver(driverID string, db *sql.DB) (journeys []*model.Journey) {
	rows, err := db.Query(`SELECT
	    j.*
	FROM
	    journeys as j,
	    drivers as d
	WHERE
	    d.id = ?
		AND j.departure_time >= d.available_from
		AND j.arrival_time <= d.available_till`, driverID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	journeys = make([]*model.Journey, 0)
	for rows.Next() {
		j := new(model.Journey)
		if err := rows.Scan(&j.ID, &j.DepartureTime, &j.ArrivalTime, &j.DepartureLocation, &j.ArrivalLocation); err != nil {
			log.Println(err)
		}
		journeys = append(journeys, j)
	}
	return
}

func FindByLocationAndTime(location string, start, end time.Time, db *sql.DB) (journeys []*model.Journey) {
	rows, err := db.Query(`SELECT
	    *
	FROM
	    journeys
	WHERE
	    departure_location = ?
		AND departure_time >= ?
		AND arrival_time <= ?`, location, start, end)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	journeys = make([]*model.Journey, 0)
	for rows.Next() {
		j := new(model.Journey)
		if err := rows.Scan(&j.ID, &j.DepartureTime, &j.ArrivalTime, &j.DepartureLocation, &j.ArrivalLocation); err != nil {
			log.Println(err)
		}
		journeys = append(journeys, j)
	}
	return
}
