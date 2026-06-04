LDFLAGS = -ldflags="-s -w"

all: build run

build:
	go build $(LDFLAGS) -o bin/goignore main.go

install:
	CGO_ENABLED=0 go install $(LDFLAGS) .

run:
	./bin/goignore
