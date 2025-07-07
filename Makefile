run:
	go run cmd/myapp/main.go

test:
	go test ./...

lint:
	golint ./...
