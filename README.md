# backend-challenge
This project is made using [Go programming language] and gorilla mux(https://go.dev/)

click [this link](https://akb6xkx9kd.execute-api.us-west-2.amazonaws.com/api/devices/20) to test the api

## Build
Run `go build -o build/main main.go` to build the project. The build artifacts will be stored in the `build/` directory.

## Test
Run `ng test ./...` to execute the unit tests 

repositories, services and also controllers have been tested.
unit tests coverage of services and controllers is 100% and unit tests coverage of repository is 90.5%;
you can see more details in cover.html file
i used **go muck** package to mock dynamodb,services and repositories.

## overview of project
in this project I used **hexagonal** architecture.this architecure is used in frameworks like nest js , asp.net ,... .in this project we have a domain package which we can define our structs and models in it.
also there is a repository which has a dependency on the database and we can interact with database in it.There are services that depend on repositories.
in services we can  validate requests and also communicate with other services or use repository to interact with detabase indirectly.each route has a specific controller.
These controlers have a dependency on services.in utils package there are some global functions and utilities and in lib package there are some functions which we can easily separate them from project
and consider as a new library.like error lib or logger lib.this project uses api gateway as proxy ,so all requests process in just one handler,and request will process by gorilla mux.
The gorilla/mux package provides request routing, validation, and other services in a straightforward, intuitive API.

## logging
In this project, zap package has been used for logging



