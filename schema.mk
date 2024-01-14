.SILENT:
.PHONY: DB_SCHEMA $(DB_SCHEMA)
.DEFAULT_GOAL := help

# commandline argument
ARG := $(word 2, $(MAKECMDGOALS))
ARG_UPPER := $(shell echo $(ARG) | tr '[:lower:]' '[:upper:]')

include ../env.sqitch
include ../../barons-court/common.mk
include ../sqitch.mk

DB_SCHEMA = $(SQITCH_SCHEMA_DEFAULT)

.PHONY: help init add deploy revert reset $(ARG) $(ARG_UPPER) $(ARG2) $(ARG2_UPPER)

init: ## init sqitch project
	$(call SQITCH-INIT)

add: ## create a sqitch migration
	$(call SQITCH-ADD)

deploy: ## deploy sqitch migrations to the database
	$(call SQITCH-DEPLOY)

revert: ## revert sqitch migrations
	$(call SQITCH-REVERT)

reset: ## reset sqitch
	$(call SQITCH-RESET)

help:
	$(call PRINT_MENU)
