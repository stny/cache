test:
	@go test

bench:
	@go test -bench=.

cover:
	@go test -cover

.PHONY: test bench cover
