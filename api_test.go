package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func init() {
	// we register an sql driver txdb
	txdb.Register("txdb", "mysql", "pie:d4e0a661552fcb11b0ba@tcp(localhost:3306)/pie?parseTime=true")
}

type apiFeature struct {
	server
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
	if a.db != nil {
		a.db.Close()
	}
	db, err := sql.Open("txdb", "api")
	if err != nil {
		panic(err)
	}
	a.db = db
}

func (a *apiFeature) iSendRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/journeys/driver/", "/journeys/driver/?driverID=DRIVER_1_ID",
		"/journeys/driver/?driverID=DRIVER_2_ID", "/journeys/driver/?driverID=DRIVER_3_ID":
		a.findByDriverHandler(a.resp, req)
	case "/journeys/location/?location=London&start=2016-03-11 03:00:00&end=2016-03-11 04:00:00",
		"/journeys/location/?location=Brighton&start=2016-03-11 09:00:00&end=2016-03-11 23:00:00",
		"/journeys/location/?location=Manchester&start=2016-03-12 03:00:00&end=2016-03-12 8:00:00":
		a.findByLocationAndTimeHandler(a.resp, req)
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		if a.resp.Code >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s",
				code, a.resp.Code, string(a.resp.Body.Bytes()))
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
	var expected, actual []byte
	var exp, act interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &exp); err != nil {
		return
	}
	if expected, err = json.MarshalIndent(exp, "", "  "); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body.Bytes(), &act); err != nil {
		return
	}
	if actual, err = json.MarshalIndent(act, "", "  "); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"expected json length: %d does not match actual: %d:\n%s",
			len(expected),
			len(actual),
			string(actual),
		)
	}

	for i, b := range actual {
		if b != expected[i] {
			return fmt.Errorf(
				"expected JSON does not match actual, showing up to last matched character:\n%s",
				string(actual[:i+1]),
			)
		}
	}
	return
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.iSendRequestTo)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)
}
