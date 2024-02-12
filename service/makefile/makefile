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
help:
	$(call PRINT_MENU, "root")
# @target$
