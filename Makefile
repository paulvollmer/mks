all: fmt lint test build

fmt:
	@go fmt

lint:
	@golint

build:
	@go build

test: test-src test-cli clean
test-src:
	@go test -cover
test-cli: build
	./mks test.txt foo
	@cat test.txt
	@echo "\n"
	./mks test/foo/bar.txt hello-world
	@cat test/foo/bar.txt
	@echo "\n"
	echo "hello stdin" | ./mks test2.txt
	@cat test2.txt
	@echo "\n"

clean:
	@rm -rf mks test*
