all: build run

build:
	go build -o bin/goignore main.go

install:
	go install -ldflags="-s -w" .

release:
	goreleaser release --auto-snapshot --clean

run:
	./bin/goignore

clean:
	rm -rf ./bin
