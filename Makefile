run :
	go run main.go

test-clean: 
	go clean -testcache

test:
	go test -v ./...
