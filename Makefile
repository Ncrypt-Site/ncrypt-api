test:
	go test -v -coverpkg=./... -coverprofile=code.coverage ./...

code-coverage: test
	go tool cover -func code.coverage