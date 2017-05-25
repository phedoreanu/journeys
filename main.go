package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olebedev/config"
	"github.com/phedoreanu/journeys/model"
)

const dateTimeForm = "2006-01-02 15:04:05"

type server struct {
	db *sql.DB
}

func (s *server) findByDriverHandler(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driverID")

	data := struct {
		Journeys []*model.Journey `json:"journeys"`
	}{Journeys: FindByDriver(driverID, s.db)}
	ok(w, data)
}

func (s *server) findByLocationAndTimeHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	start, _ := time.Parse(dateTimeForm, r.URL.Query().Get("start"))
	end, _ := time.Parse(dateTimeForm, r.URL.Query().Get("end"))

	data := struct {
		Journeys []*model.Journey `json:"journeys"`
	}{Journeys: FindByLocationAndTime(location, start, end, s.db)}
	ok(w, data)
}

func main() {
	cfg, err := config.ParseYamlFile("config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	activeProfile := os.Getenv("ACTIVE_PROFILE")
	if activeProfile == "" {
		activeProfile = "dev"
	}

	cfg, err = cfg.Get(activeProfile)
	if err != nil {
		log.Fatalln(err)
	}

	dbURI, err := cfg.String("dbURI")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("Staring Journeys API:4000")

	s := &server{db: db}
	http.HandleFunc("/journeys/driver/", s.findByDriverHandler)
	http.ListenAndServe(":4000", nil)
}

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(resp))
}

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if s, ok := data.(string); ok {
		fmt.Fprintf(w, s)
		return
	}

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fail(w, "oops something evil has happened", 500)
		return
	}

	fmt.Fprintf(w, string(resp))
}
