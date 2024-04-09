build:
	@go build -o bin/coinmarketcap
run: build
	@./bin/coinmarketcap
test:
	@go test -v ./...