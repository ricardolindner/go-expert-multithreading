# go-expert-multithreading
This project is a challenge from the **`Go Expert`** course, focused on solving problems using **multithreading** and **HTTP requests** in Go.

The goal is to fetch address data by CEP from two distinct APIs **simultaneously**, and return the **fastest** response.

## Download Dependencies
From the root directory, run:
```bash
go mod tidy
```

## Run the Server
Navigate to the server folder and start the application:
```bash
cd cmd/server
go run main.go
```

## Sample Request
With the server running locally, execute a `GET` request for the url `http://localhost:8000/cep/{CEP}`:
```bash
curl http://localhost:8000/cep/70040010
```

## Project Structure
```text
go-expert-multithreading/
|-- cmd/
|   |-- server/     # Application entry point
|       |-- main.go
|-- internal/
|   |-- infra/
|       |-- webserver/
|           |-- handlers/ # Handlers HTTP
|-- go.mod
|-- README.md
```

## Technologies
* Go 1.23+
* net/http
* context.WithTimeout
* Goroutines and channels