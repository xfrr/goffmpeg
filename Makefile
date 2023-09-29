

test:
	go test -v ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./...

coverage-html:
	go tool cover -html=coverage.out