.PHONY: run
run:
	go run main/main.go

.PHONY: test
test:
	go test -v ./test