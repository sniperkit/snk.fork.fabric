default: build test

build:
	go install -v ./...

get: 
	go get ./...

repo: get build

test: 
	go test -v -tags test ./...