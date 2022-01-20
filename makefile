SHELL := /bin/bash

# Modules support

tidy:
	go mod tidy
	go mod vendor

# to lunch: make tidy

run:
go run main.go