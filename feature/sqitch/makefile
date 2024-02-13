# @include@
include env.sqitch
include ./.makefile/postgres.mk
include ./.makefile/sqitch.mk
# @include$

# @include-post@
DB_SCHEMA = $(DB_SCHEMA_DEFAULT)

.PHONY: help create drop create-schema remove-schema $(ARG)
# @include-post$

# @target@
create: ## create database
	$(call DB-CREATE)

drop: ## drop database
	$(call DB-DROP)

create-schema: ## create sqitch project (schema)
	$(call DB-CREATE-SCHEMA)

remove-schema: ## remove sqitch project (directories and meta-data)
	$(call DB-DROP-SCHEMA)
	$(call DB-DROP-SCHEMA-META)
	$(call DB-REMOVE-SCHEMA)

setup: ## setup sqitch locally
	brew tap sqitchers/sqitch
	brew install sqitch --with-postgres-support	

# @target$
