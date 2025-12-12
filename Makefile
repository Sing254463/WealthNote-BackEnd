# Makefile

.PHONY: run build clean

run:
	go run cmd/main.go

build:
	go build -o WealthNoteBackend cmd/main.go

clean:
	go clean
	rm -f WealthNoteBackend

install:
	go get ./...