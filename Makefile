build:
	go build -o qsoft-test ./cmd/main.go
run:
	./qsoft-test
.DEFAULT_GOAL := build