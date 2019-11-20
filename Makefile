all: fmt lint test

fmt:
	@go fmt

lint:
	@golint

build:
	@go build

test: build
	./mkt test.txt foo
	@echo ""
	cat test.txt
	@echo "\n--------------------------------"
	./mkt test/foo/bar.txt hello-world
	@echo ""
	cat test/foo/bar.txt
	@echo "\n--------------------------------"

clean:
	@rm -rf mkt test*
