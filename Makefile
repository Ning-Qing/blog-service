.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	swag init
	go build -o build/example main.go
	