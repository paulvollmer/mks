all: lint test build

lint-prepare:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2
lint:
	@./bin/golangci-lint run ./...

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
