
compile: clean
	@go build -o build/tangram

clean:
	@rm -rf build/

build: 
	@CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" -a -installsuffix cgo -o build/tangram .

