# Journeys
A small micro-service designed to find journeys in its MySQL database.

## Testing
Install `godog`: 
```bash
$ go get -u github.com/DATA-DOG/godog/cmd/godog 
```
and test the BDDs:
```bash
$ godog find_by_driver.feature
$ godog find_by_location_and_time.feature
```

## Deployment
To build the Docker image run: 
```bash
$ ./bin/docker
```
and deploy the resulting image i.e. `phedoreanu/journeys:1.0.0`.
