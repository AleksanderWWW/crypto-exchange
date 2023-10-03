build:
	go build -o ./bin/exchange

run: build
	./bin/exchange

test:
	go test -v -cover ./...

fmt:
	go fmt ./...
