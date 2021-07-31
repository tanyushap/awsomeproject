BUILD_DIR := bin

.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test -v ./test