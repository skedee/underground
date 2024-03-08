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
docs-serve: ## serve docsify
	echo "Starting docsify"
	docsify serve .

docs-setup: ## setup docsify
	@echo "npm i docsify-cli"
	sudo npm i docsify-cli -g
	
help:
	$(call PRINT_MENU, "root")
# @target$
