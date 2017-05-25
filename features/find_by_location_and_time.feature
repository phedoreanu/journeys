# file: find_by_location_and_time.feature
Feature: get all the journeys by location and time
  A journey must start at X
  And overlap with the given time range

  Scenario: No journeys from London on 2016-03-11 between 3am and 4am
    When I send "GET" request to "/journeys/location/?location=London&start=2016-03-11 03:00:00&end=2016-03-11 04:00:00"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": []
      }
      """

  Scenario: 2 journeys from Brighton on 2016-03-11 between 9am and 23pm
    When I send "GET" request to "/journeys/location/?location=Brighton&start=2016-03-11 09:00:00&end=2016-03-11 23:00:00"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": [
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-11T20:00:00Z",
            "departure_location": "Brighton",
            "departure_time": "2016-03-11T18:00:00Z",
            "id": "ROUTE_2_ID"
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

  Scenario: 1 journey from Manchester on 2016-03-11 between 3am and 8am
    When I send "GET" request to "/journeys/location/?location=Manchester&start=2016-03-12 03:00:00&end=2016-03-12 8:00:00"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "journeys": [
          {
            "arrival_location": "London",
            "arrival_time": "2016-03-12T07:00:00Z",
            "departure_location": "Manchester",
            "departure_time": "2016-03-12T03:00:00Z",
            "id": "ROUTE_10_ID"
          }
        ]
      }
      """
