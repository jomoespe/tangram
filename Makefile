
compile: clean
	@go build -o build/tangram

clean:
	@rm -rf build/

build: 
	@CGO_ENABLED=0 GOOS=linux go build -o build/tangram -a -installsuffix cgo

