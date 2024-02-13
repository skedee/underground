# @include-pre@
GO_VERSION = 1.22.0

# @include-pre$

# @target@
build: ## build go service
	go build

install: ## install go service
	go install

setup-go: ## setup golang
	brew install wget
	wget https://go.dev/dl/go$(GO_VERSION).darwin-arm64.pkg
	sudo installer -pkg go$(GO_VERSION).darwin-arm64.pkg -target /Applications
	rm go$(GO_VERSION).darwin-arm64.pkg

# @target$
