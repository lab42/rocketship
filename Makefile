help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: ## Initialize required git hooks
	git config core.hooksPath ./.githooks

dev: ## Development (Runs all test suites on .go file changes)
	find **/*.go | entr sh -c '$$(go env GOPATH)/bin/ginkgo --succinct ./...'

test: ## Run all test suites
	$$(go env GOPATH)/bin/ginkgo --succinct ./...

build: ## Build binaries
	# CGO_ENABLED=0 go build -a -ldflags '-linkmode external -extldflags "-static" -w' -v
	make build_all -j 3

# Private targets
build_all: build_linux build_darwin build_windows

build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/linux/rocketship

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/darwin/rocketship

build_windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/windowd/rocketship.exe
