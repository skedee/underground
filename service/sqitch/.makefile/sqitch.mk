# Read the number from the version file or set a default value of 0
CURRENT_NUMBER := $(shell if [ -e version ]; then cat version; else echo 0; fi)

# Increment the version number
NEW_NUMBER := $(shell expr $(CURRENT_NUMBER) + 1)

# Format the version number with a fixed width (000)
MIGRATION_FILE := $(shell printf "%03d-%s" $(NEW_NUMBER) $(ARG_UPPER))

DB_SCHEMA := $(shell basename `pwd`)

ifdef ARG_UPPER
	ifneq ($(wildcard deploy),)
		CURRENT_MIGRATION_FILE := $(shell printf "deploy/%03d-%s" $(CURRENT_NUMBER) $(ARG_UPPER))
		NEWEST_FILE := $(shell find deploy -type f -name "*$(ARG_UPPER).sql" -exec ls -lt {} + | awk 'NR==1 {print $$9}')
	endif
endif

# init the sqitch project (directories and meta-data)
define SQITCH-INIT
	sqitch init $(DB_NAME)
	echo "sqitch init $(DB_NAME)"
endef

# add a new migration. copy the previous content if to new migration (if present).
define SQITCH-ADD
	sqitch add $(MIGRATION_FILE)
	@if [ "$(NEWEST_FILE)" = "" ]; then \
		echo "$(CURRENT_MIGRATION_FILE) does not exist"; \
	else \
		cp -f $(NEWEST_FILE) deploy/$(MIGRATION_FILE).sql; \
	fi
	echo $(NEW_NUMBER) > version;
endef

# delete all sqitch files except for the sqitch.conf file.
define SQITCH-RESET
	rm -rf deploy revert verify version sqitch.plan
	echo "sqitch reset"
endef

# deploy sqitch migrations to the database
define SQITCH-DEPLOY
	sqitch deploy
endef

# revert sqitch changes
define SQITCH-REVERT
	sqitch revert
endef

# remove the sqitch meta database
define SQITCH-REMOVE
	sqitch remove db:pg://$(DB_USER):$(PGPASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)
endef

# remove the database
define SQITCH-DROP
	psql -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS DROP $(DB_SCHEMA);"
endef


# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest