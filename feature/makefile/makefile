# @include-pre@
.SILENT:
.DEFAULT_GOAL := help
.PHONY: docs docs-setup help

OS = $(shell uname -s)

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
# @include-post$

# @target@
docs: ## serve docsify
	@docsify serve . && \
	{ \
		printf "❌ $(RED)Error Starting docsify$(NC)\n"; \
		printf "	🔹 $(YELLOW)Try running `docs-setup`.$(NC)\n"; \
	}

docs-setup: ## setup docsify
	@echo "npm i docsify-cli"
	sudo npm i docsify-cli -g
	
about: ## display about
	@sed -n '1p' README.md | tr -d '#'
		
help:
	$(call PRINT_MENU, "root")
# @target$
