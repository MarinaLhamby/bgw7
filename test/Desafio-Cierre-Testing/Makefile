test:
	- go test ./... -count=1

coverage: 
	- go test -coverprofile=coverage.out ./...

html-report: coverage
	- go tool cover -html=coverage.out

mock:
	- go generate ./...
