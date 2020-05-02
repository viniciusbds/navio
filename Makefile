all: build

build:
	go build -o navio  main.go

run:
	./navio

clean:
	rm navio

unit-tests:
	go test ./assert/...
	go test ./container/...
	go test ./images/...
	go test ./logger/...
	go test ./utilities/...