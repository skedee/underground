# @include-pre@
.PHONY: generate 
GO_VERSION = 1.22.2

.PHONY: build install tag tag-latest run watch test test-all

.SILENT:
APP_NAME=driskill
.DEFAULT_GOAL := help
.PHONY: docs docs-setup help

# commandline argument
ARG := $(word 2, $(MAKECMDGOALS))
ARG_UPPER := $(shell echo $(ARG) | tr '[:lower:]' '[:upper:]')

# domain name
DOMAIN=$$(git config user.email | sed 's/.*@//; s/\..*//')
# @include-pre$

# @include@
include ./.makefile/common.mk
# @include$

# @include-post@
# hostname - used to create swagger url
ifndef APP_URL
	HOSTNAME="localhost"
	HOST_PORT=8080
else
	HOSTNAME=$(APP_URL)
	HOST_PORT=$(PORT)
endif
# @include-post$

# @target@
swag: ## run swag init
	@swag init > /dev/null 2>&1 || (echo "❌ Error: failed to initialize swagger"; swag init; exit 1)
	@if [ "$(OS)" = "Darwin" ]; then \
		sed -i '' "s/@@HOSTNAME/$(HOSTNAME)"/g docs/docs.go; \
		sed -i '' "s/@@HOSTNAME/$(HOSTNAME)"/g docs/swagger.json; \
		sed -i '' "s/@@HOSTNAME/$(HOSTNAME)"/g docs/swagger.yaml; \
		sed -i '' "s/@@PORT/$(HOST_PORT)"/g docs/docs.go; \
		sed -i '' "s/@@PORT/$(HOST_PORT)"/g docs/swagger.json; \
		sed -i '' "s/@@PORT/$(HOST_PORT)"/g docs/swagger.yaml; \
		echo "✅ swagger generate successful"; \
	else \
		sed -i "s/@@HOSTNAME/$(HOSTNAME)"/g docs/docs.go; \
		sed -i "s/@@HOSTNAME/$(HOSTNAME)"/g docs/swagger.json; \
		sed -i "s/@@HOSTNAME/$(HOSTNAME)"/g docs/swagger.yaml; \
		sed -i "s/@@PORT/$(HOST_PORT)"/g docs/docs.go; \
		sed -i "s/@@PORT/$(HOST_PORT)"/g docs/swagger.json; \
		sed -i "s/@@PORT/$(HOST_PORT)"/g docs/swagger.yaml; \
		echo "✅ swagger generate successful"; \
	fi

build-debug: generate swag ## build go service
	rm -f __debug*
	go mod tidy
	@go build -gcflags "all=-N -l"  -o ${APP_NAME} || (echo "❌ Error: Failed to build the application"; exit 1)
	@echo "✅ build successful"

build: generate swag ## build go service
	rm -f __debug*
	go mod tidy
	@go build -o ${APP_NAME} || (echo "❌ Error: Failed to build the application"; exit 1)
	@echo "✅ build successful"

build-linux-amd64: generate ## build go service for linux amd64
	go mod tidy
	@GOOS=linux GOARCH=amd64 go build -o ${APP_NAME} || (echo "❌ Error: Failed to build the application"; exit 1)
	@echo "✅ build successful"

build-linux-arm64: generate ## build go service for linux arm64
	go mod tidy
	@GOOS=linux GOARCH=arm64 go build -o ${APP_NAME} || (echo "❌ Error: Failed to build the application"; exit 1)
	@echo "✅ build successful"

run: build ## run the go service
	@if pgrep -x "$(APP_NAME)" > /dev/null; then \
		pkill $(APP_NAME); \
	fi
	./$(APP_NAME) &

watch: ## watch and rebuild on file changes
	echo "Watching for changes to *.go files..."
	$(MAKE) run
	@while true; do \
		trap 'break' INT; \
		fswatch -r -1 src; \
		$(MAKE) run; \
	done

install: ## install go service
	@go install || (echo "❌ Error: Failed to install"; exit 1)
	printf "✅ $(GREEN)install successful$(NC)\n" \

kill: ## kills service
	-pkill $(APP_NAME) >/dev/null || (echo "ℹ️  Info: $(APP_NAME) not running"; exit 1)
	printf "✅ $(GREEN)$(APP_NAME)$(NC)\n"

tag-latest: ## get the latest github tag for the go service
	git ls-remote --tags --refs https://github.com/$(DOMAIN)/$(APP_NAME) | awk '{print $2}' | cut -d '/' -f 3 | sort -V | tail -n 1

tag: ## create a new github tag for go service
ifdef ARG
	@git tag -a $(ARG) -m '$(ARG)' || (echo "❌ Error: Failed to create tag"; exit 1)
	@git push origin $(ARG) || (echo "❌ Error: Failed push tag"; exit 1)
	printf "✅ $(GREEN)tag created successful$(NC)\n" \
else
	echo "❌ Error: tag argument missing. i.e. v0.0.1"
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

go-update: ## go update dependencies
	@go get -u || (echo "❌ Error: Failed to update dependencies"; exit 1)
	printf "✅ $(GREEN)update successful$(NC)\n" \

go-setup: ## install or upgrade golang
	@if [ "$(OS)" = "Darwin" ]; then \
		brew install wget > /dev/null 2>&1 || (echo "❌ Error: Failed to install wget"; exit 1); \
		printf "✅ $(GREEN)wget installed successful$(NC)\n"; \
		brew install fswatch > /dev/null 2>&1 || (echo "❌ Error: Failed to install fswatch"; exit 1); \
		printf "✅ $(GREEN)fswatch installed successful$(NC)\n"; \
		wget https://go.dev/dl/go$(GO_VERSION).darwin-arm64.pkg > /dev/null 2>&1 || (echo "❌ Error: Failed to get go$(GO_VERSION).darwin-arm64.pkg"; exit 1); \
		sudo installer -pkg go$(GO_VERSION).darwin-arm64.pkg -target /Applications > /dev/null 2>&1 || (echo "❌ Error: Failed to install go$(GO_VERSION)"; exit 1); \
		printf "✅ $(GREEN)go$(GO_VERSION).darwin-arm64.pkg installed $(NC)\n"; \
		rm go$(GO_VERSION).darwin-arm64.pkg* > /dev/null 2>&1; \
	elif [ "$(IS_ARM)" = "true" ]; then \
		wget https://go.dev/dl/go$(GO_VERSION).linux-arm64.tar.gz > /dev/null 2>&1 || (echo "❌ Error: Failed to get go$(GO_VERSION).linux-arm64.tar.gz"; exit 1); \
		sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-arm64.tar.gz > /dev/null 2>&1 || (echo "❌ Error: Failed to install go$(GO_VERSION).linux-arm64.tar.gz"; exit 1); \
		printf "✅ $(GREEN)go$(GO_VERSION).linux-arm64.tar.gz installed $(NC)\n"; \
	else \
		wget https://go.dev/dl/go$(GO_VERSION).linux-amd64.tar.gz > /dev/null 2>&1 || (echo "❌ Error: Failed to get go$(GO_VERSION).linux-amd64.tar.gz"; exit 1); \
		sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz > /dev/null 2>&1 || (echo "❌ Error: Failed to install go$(GO_VERSION).linux-amd64.tar.gz"; exit 1); \
		printf "✅ $(GREEN)go$(GO_VERSION).linux-amd64.tar.gz installed $(NC)\n"; \
	fi
	printf "✅ $(GREEN)update ~/.bashrc with (GOPATH, PATH)$(NC)\n"
	printf '	$(BLUE)export GOPATH=$$(HOME)/go$(NC)\n'
	printf '	$(BLUE)export PATH=$$PATH:/usr/local/go/bin:$$GOPATH/bin$(NC)\n'
	@go install github.com/swaggo/swag/cmd/swag@latest > /dev/null 2>&1 || (echo "❌ Error: Failed to install swag"; exit 1)
	printf "✅ $(GREEN)swag installed successful$(NC)\n"

help:
	$(call PRINT_MENU, "root")
# @target$
