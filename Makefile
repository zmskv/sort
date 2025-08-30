
.PHONY: build
build:
	go build -o sort cmd/sort/main.go

.PHONY: install
install:
	go mod download

.PHONY : clear
clear:
	go mod tidy

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: test100
test100:
	go test -v -count=100 ./...

.PHONY: race
race:
	go test -v -race -count=1 ./...
