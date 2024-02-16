# @include-pre@
GO_VERSION = 1.22.0

# @include-pre$

# @target@
build: ## build go service
	go build

install: ## install go service
	go install

tag-latest: ## get the latest github tag for the go service
	git ls-remote --tags --refs https://github.com/@@project-name@@ | awk '{print $2}' | cut -d '/' -f 3 | sort -V | tail -n 1
	
tag: ## create a new github tag for go service
ifdef ARG
	git tag -a $(ARG) -m '$(ARG)'
	git push origin $(ARG)
else
	echo "github: tag argument missing. i.e. v0.0.1"
endif
		
test: build ## run unit test in current directory. v for verbose
ifdef ARG	
	go test -v
else
	go test
endif

test-all: build ## run all unit tests in project. v for verbose
ifdef ARG
	go test ./... -v
else
	go test ./...
endif

setup-go: ## install or upgrade golang
	@if [ "$(OS)" = "Darwin" ]; then \
		brew install wget; \
		wget https://go.dev/dl/go$(GO_VERSION).darwin-arm64.pkg; \
		sudo installer -pkg go$(GO_VERSION).darwin-arm64.pkg -target /Applications; \
		rm go$(GO_VERSION).darwin-arm64.pkg; \
	fi

# @target$
