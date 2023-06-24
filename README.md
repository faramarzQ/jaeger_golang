# jeager_golang
A simple Jaeger open-tracing project implemented on golang

This project demonstrate a simple usage of Jaeger in a golang applicatiion implemented to cover the contents of [this article](https://sokanacademy.com/plus/faramarzq/jaeger-introduction-with-golang). 

## Installation


- Run Jaeger server using:    

>    `make up`    

- Download project's dependencies:

>    `go mod download`

- Run server application using:    

>    `make run-app`    

- open up [http://localhost:8080](http://localhost:8080) to send a sample request
