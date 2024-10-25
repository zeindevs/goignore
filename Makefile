run: build
	./bin/goignore

build:
	go build -o bin/goignore main.go

install:
	go install -ldflags="-s -w" .
