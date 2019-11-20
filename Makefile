all: fmt lint test

fmt:
	@go fmt

lint:
	@golint

build:
	@go build

test: build
	./mks test.txt foo
	@echo ""
	cat test.txt
	@echo "\n--------------------------------"
	./mks test/foo/bar.txt hello-world
	@echo ""
	cat test/foo/bar.txt
	@echo "\n--------------------------------"

clean:
	@rm -rf mks test*
