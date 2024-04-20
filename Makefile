.PHONY: test-lexer run

run:
	@go run main.go

test-lexer:
	@go test ./lexer -v
