# @include-pre@
GO_VERSION = 1.22.1

.PHONY: build install tag tag-latest run watch test test-all

# @include-pre$

# @target@
build-darwin: generate ## build go service
	go mod tidy
	swag init
	GOOS=darwin GOARCH=amd64 go build -o ${APP_NAME}

build-linux: generate ## build go service
	go mod tidy
	swag init	
	GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}	

run: build-darwin ## run the go service
	@if pgrep -x "$(APP_NAME)" > /dev/null; then \
		pkill $(APP_NAME); \
	fi
	./$(APP_NAME) &

watch: ## watch and rebuild on file changes
	echo "Watching for changes to *.go files..."
	$(MAKE) run
	@while true; do \
		trap 'break' INT; \
		fswatch -r -1 api bo dao shared sqlc/query validate; \
		$(MAKE) run; \
	done

install: ## install go service
	go install

tag-latest: ## get the latest github tag for the go service
	git ls-remote --tags --refs https://github.com/$(DOMAIN)/$(APP_NAME) | awk '{print $2}' | cut -d '/' -f 3 | sort -V | tail -n 1
	
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
		brew install fswatch; \
		wget https://go.dev/dl/go$(GO_VERSION).darwin-arm64.pkg; \
		sudo installer -pkg go$(GO_VERSION).darwin-arm64.pkg -target /Applications; \
		rm go$(GO_VERSION).darwin-arm64.pkg; \
	else \
		wget https://go.dev/dl/go$(GO_VERSION).linux-amd64.tar.gz; \
		sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz; \
		echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc; \
		echo 'export GOPATH=$HOME/go' >> ~/.bashrc; \
		echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc; \
		source ~/.bashrc; \		
	fi
	go install github.com/swaggo/swag/cmd/swag@latest

# @target$
