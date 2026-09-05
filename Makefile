VERSION = $(shell git describe --tags --abbrev=0)
LDFLAGS = -ldflags="-s -w"

all: build run

build:
	go build $(LDFLAGS) -o bin/goignore .

install:
	go install $(LDFLAGS) .

release:
	goreleaser release --auto-snapshot --clean

run:
	./bin/goignore

clean:
	rm -rf ./bin
