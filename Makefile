all: build

build:
	@echo "Building..."	
	@go build -o navio  main.go
	@echo "ok"

run:
	@./navio

clean:
	@rm navio

unit-tests:
	@go test ./container/... ./images/... ./logger/... ./utilities/...