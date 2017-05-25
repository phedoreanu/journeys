# Journeys
A small RESTful micro-service designed to find journeys in its MySQL database.

## Testing
Install `godog`: 
```bash
$ go get -u github.com/DATA-DOG/godog/cmd/godog 
```
and test the BDDs:
```bash
$ ./bin/bdds
```

## Deployment
To build the Docker image run: 
```bash
$ ./bin/docker
```
and deploy the resulting image i.e. `phedoreanu/journeys:1.0.0`.

## Implementations details
This API has 2 RESTFul endpoints:
 * GET /journeys/driver/?driverID={driverID}
 * GET /journeys/location/?location={location}&start={start}&end={end}
 
Please check the [features](https://github.com/phedoreanu/journeys/tree/master/features) folder for example responses. 

_Note: Both {start} and {end} use this datetime format `2016-03-11 03:00:00`_

## Explanations
 * Using `github.com/olebedev/config` for configuration based on ENV vars i.e `ACTIVE_PROFILE=production`
 * Docker script runs the BDDs before building the image
 