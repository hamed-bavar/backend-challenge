# backend-challenge
This project is made using [Go programming language](https://go.dev/)  and gorilla mux

click [this link](https://akb6xkx9kd.execute-api.us-west-2.amazonaws.com/api/devices/20) to test the api

## Build
Run `go build -o build/main main.go` to build the project. The build artifacts will be stored in the `build/` directory.

## Test
Run `ng test ./...` to execute the unit tests 

repositories, services and also controllers have been tested.
unit tests coverage of services and controllers is 100% and unit tests coverage of repository is 90.5%.
you can see more details in cover.html file.
I used **go muck** package to mock dynamodb,services and repositories.

## overview of project structure
In this project, I used hexagonal architecture. This architecture is utilized in frameworks like NestJS, ASP.NET, and others. In this project, we have a domain package where we can define our structs and models (and also our repository interface). Additionally, there is a repository that depends on the database, enabling us to interact with it. The services rely on the repositories.

Within the services, we can validate requests, communicate with other services, or use the repository to interact with the database indirectly. Each route has a specific controller. These controllers have a dependency on services.

In the utils package, there are some global functions and utilities, while in the lib package, there are functions that can be easily separated from the project and treated as a new library (e.g., error lib or logger lib). For proxy, this project uses API gateway, capturing all requests by one lambda, and gorilla mux processes these requests.

The gorilla/mux package provides request routing, validation, and other services in a straightforward, intuitive API.

## logging
In this project, zap package has been used for logging



