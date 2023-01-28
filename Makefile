.PHONY: clean
clean:
	\rm -rf bin/*

build: clean
	go build -o ./bin/citypair ./cmd/main.go

run:
	go run ./cmd/main.go

start:
	./bin/citypair

build-all: clean
	GOOS=linux GOARCH=amd64 go build -o ./bin/citypair-linux-amd64 ./cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o ./bin/citypair-linux-arm64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o ./bin/citypair-darwin-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./bin/citypair-darwin-arm64 ./cmd/main.go

deps-cleancache:
	go clean -modcache

mock:
	mockgen -source=internal/flight/service.go \
		-package testutil \
		-destination=testutil/mocks/flight/service.go