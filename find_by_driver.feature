# file: find_by_driver.feature
Feature: get all the journeys by driver
  In order to drive
  A driver must be available
  The whole journey

  Scenario: DRIVER_0_ID doesn't exist
    When I send "GET" request to "/journeys/driver/"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": []
      }
      """

  Scenario: DRIVER_1_ID available on 2016-03-11 from 9am to 23pm
    When I send "GET" request to "/journeys/driver/?driverID=DRIVER_1_ID"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": [
          {
            "arrival_location": "Birmingham",
            "arrival_time": "2016-03-11T23:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T20:00:00Z",
            "id": "ROUTE_11_ID"
          },
          {
            "arrival_location": "Brighton",
            "arrival_time": "2016-03-11T17:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T15:00:00Z",
            "id": "ROUTE_1_ID"
          },
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T20:00:00Z",
            "departure_location": "Brighton",
            "departure_time": "2016-03-11T18:00:00Z",
            "id": "ROUTE_2_ID"
          },
          {
            "arrival_location": "Manchester",
            "arrival_time": "2016-03-11T13:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T09:00:00Z",
            "id": "ROUTE_3_ID"
          },
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T18:00:00Z",
            "departure_location": "Manchester",
            "departure_time": "2016-03-11T14:00:00Z",
            "id": "ROUTE_4_ID"
          },
          {
            "arrival_location": "Birmingham",
            "arrival_time": "2016-03-11T15:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T12:00:00Z",
            "id": "ROUTE_5_ID"
          },
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T19:00:00Z",
            "departure_location": "Birmingham",
            "departure_time": "2016-03-11T16:00:00Z",
            "id": "ROUTE_6_ID"
          },
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T11:00:00Z",
            "departure_location": "Brighton",
            "departure_time": "2016-03-11T09:00:00Z",
            "id": "ROUTE_8_ID"
          }
        ]
      }
      """

  Scenario: DRIVER_2_ID available on 2016-03-11 from 12am to 6pm
    When I send "GET" request to "/journeys/driver/?driverID=DRIVER_2_ID"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": [{
            "id": "ROUTE_1_ID",
            "departure_time": "2016-03-11T15:00:00Z",
            "departure_location": "London",
            "arrival_time": "2016-03-11T17:00:00Z",
            "arrival_location": "Brighton"
        }, {
            "id": "ROUTE_4_ID",
            "departure_time": "2016-03-11T14:00:00Z",
            "departure_location": "Manchester",
            "arrival_time": "2016-03-11T18:00:00Z",
            "arrival_location": "London"
        }, {
            "id": "ROUTE_5_ID",
            "departure_time": "2016-03-11T12:00:00Z",
            "departure_location": "London",
            "arrival_time": "2016-03-11T15:00:00Z",
            "arrival_location": "Birmingham"
        }]
      }
      """

  Scenario: DRIVER_3_ID available on 2016-03-11 from 9am to 5pm
    When I send "GET" request to "/journeys/driver/?driverID=DRIVER_3_ID"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": [
          {
            "arrival_location": "Brighton",
            "arrival_time": "2016-03-11T17:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T15:00:00Z",
            "id": "ROUTE_1_ID"
          },
          {
            "arrival_location": "Manchester",
            "arrival_time": "2016-03-11T13:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T09:00:00Z",
            "id": "ROUTE_3_ID"
          },
          {
            "arrival_location": "Birmingham",
            "arrival_time": "2016-03-11T15:00:00Z",
            "departure_location": "London",
            "departure_time": "2016-03-11T12:00:00Z",
            "id": "ROUTE_5_ID"
          },
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T11:00:00Z",
            "departure_location": "Brighton",
            "departure_time": "2016-03-11T09:00:00Z",
            "id": "ROUTE_8_ID"
          }
        ]
      }
      """
