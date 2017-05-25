package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/phedoreanu/journeys/model"
)

// FindByDriverID returns a array of journeys available
// for the given driverID
func FindByDriverID(driverID string, db *sql.DB) []*model.Journey {
	return queryJourneys(db, `SELECT
	    j.*
	FROM
	    journeys as j,
	    drivers as d
	WHERE
	    d.id = ?
		AND j.departure_time >= d.available_from
		AND j.arrival_time <= d.available_till`, driverID)
}

// FindByLocationAndTime returns a array of journeys
// from the given location and overlap with the given time frame
func FindByLocationAndTime(location string, start, end time.Time, db *sql.DB) []*model.Journey {
	return queryJourneys(db, `SELECT
	    *
	FROM
	    journeys
	WHERE
	    departure_location = ?
		AND departure_time >= ?
		AND arrival_time <= ?`, location, start, end)
}

func queryJourneys(db *sql.DB, query string, args ...interface{}) (journeys []*model.Journey) {
	rows, err := db.Query(query, args...)
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
