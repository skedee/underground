# @include-pre@
.SILENT:
.DEFAULT_GOAL := help

# commandline argument
ARG := $(word 2, $(MAKECMDGOALS))
ARG_UPPER := $(shell echo $(ARG) | tr '[:lower:]' '[:upper:]')
# @include-pre$

# @include@
include ./.makefile/common.mk
# @include$

# @include-post@
# @include-post$

# @target@
docs-serve: ## serve docsify
	docsify serve .

docs-setup: ## setup docsify
	@echo "npm i docsify-cli"
	sudo npm i docsify-cli -g
	
help:
	$(call PRINT_MENU, "root")
# @target$
