# @include-pre@
GO_VERSION = 1.22.0
OS = $(shell uname -s)

# @include-pre$

# @target@
build: ## build go service
	go build

install: ## install go service
	go install

setup: ## install or upgrade golang
	@if [ "$(OS)" = "Darwin" ]; then \
		brew install wget; \
		wget https://go.dev/dl/go$(GO_VERSION).darwin-arm64.pkg; \
		sudo installer -pkg go$(GO_VERSION).darwin-arm64.pkg -target /Applications; \
		rm go$(GO_VERSION).darwin-arm64.pkg; \
	fi

# @target$
