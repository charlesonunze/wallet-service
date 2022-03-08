run:
	go run .

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

buf:
	buf mod update && buf generate

.PHONY: run sqlc test buf