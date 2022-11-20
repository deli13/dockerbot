export GO111MODULE=on

deps:

run:
	go run main.go

build:
	go mod download && CGO_ENABLED=0 go build -o ./bin/bot ./main.go
