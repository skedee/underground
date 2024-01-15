.SILENT:
.DEFAULT_GOAL := help

# commandline argument
ARG := $(word 2, $(MAKECMDGOALS))
ARG_UPPER := $(shell echo $(ARG) | tr '[:lower:]' '[:upper:]')

include env.sqitch
include ../barons-court/common.mk
include postgres.mk
include sqitch.mk

DB_SCHEMA = $(DB_SCHEMA_DEFAULT)

.PHONY: help create drop create-schema remove-schema $(ARG)

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

help:
	$(call PRINT_MENU, "root")
