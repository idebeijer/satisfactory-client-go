.PHONY: test test-short

# Command to run all tests
test:
	go test ./...

# Command to run tests in short mode
test-short:
	go test -short ./...