test:
	go test -coverprofile=coverage.out -v $$(go list ./... | grep -v fixtures)

watch:
	@go run main.go watch

run:
	@go run main.go run