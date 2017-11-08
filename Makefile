
compile: clean
	@go build 

clean:
	@go clean

dependencies:
	@dep ensure

build: 
	@CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X 'main.version=$(shell cat VERSION)' -X 'main.build=$(shell git rev-parse --short HEAD)' -X 'main.buildDate=$(shell date --rfc-3339=seconds)'" -a -installsuffix cgo
