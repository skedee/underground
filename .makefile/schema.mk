.SILENT:
.PHONY: DB_SCHEMA $(DB_SCHEMA)
.DEFAULT_GOAL := help

# commandline argument
ARG := $(word 2, $(MAKECMDGOALS))
ARG_UPPER := $(shell echo $(ARG) | tr '[:lower:]' '[:upper:]')

include ../env.sqitch
include ../../barons-court/common.mk
include ../.makefile/sqitch.mk
include ../.makefile/postgres.mk

DB_SCHEMA = $(SQITCH_SCHEMA_DEFAULT)

.PHONY: help init add deploy revert reset $(ARG) $(ARG_UPPER) $(ARG2) $(ARG2_UPPER)

init: ## init the sqitch project (directories and meta-data)
	$(call SQITCH-INIT)

add: ## create a sqitch migration
	$(call SQITCH-ADD)

deploy: ## deploy sqitch migrations to the database
	$(call SQITCH-DEPLOY)

revert: ## revert sqitch migrations
	$(call SQITCH-REVERT)

backup: ## backup schema
	$(call DB-BACKUP-SCHEMA, $(shell basename `pwd`))

restore: ## restore schema
	$(call DB-RESTORE-SCHEMA, $(shell basename `pwd`))

generate: ## generate queries
	sqlc generate

remove-schema: ## remove sqitch project
	$(call DB-DROP-SCHEMA, $(shell basename `pwd`))
	$(call DB-DROP-SCHEMA-META, $(shell basename `pwd`))

remove-schema-meta: ## remove sqitch project (meta-data)
	$(call DB-DROP-SCHEMA-META, $(shell basename `pwd`))

backup-schema-meta: ## backup schema meta-data
	$(call DB-BACKUP-SCHEMA-META, $(shell basename `pwd`))

restore-schema-meta: ## restore schema meta-data
	$(call DB-RESTORE-SCHEMA-META, $(shell basename `pwd`))

reset: ## reset sqitch
	$(call SQITCH-RESET)

help:
	$(call PRINT_MENU)
