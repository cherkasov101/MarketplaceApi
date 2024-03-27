FILES = ./cmd/web

all: clean build
	./main

build:
	go build -o main ./cmd/web

fmt:
	gofmt -w .

clean:
	rm -rf main