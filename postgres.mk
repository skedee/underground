.SILENT:

db-create-schema:
	cp sqitch.conf $(ARG);
	cp schema.mk $(ARG)/makefile;
	sed -i "s/@@registry/registry = sqitch_$(ARG)/g" $(ARG)/sqitch.conf
	sed -i "s/@@db_name/$(DB_NAME)"/g $(ARG)/sqitch.conf
	sed -i "s/@@sqitch_user/$(SQITCH_USER)"/g $(ARG)/sqitch.conf
	sed -i "s/@@sqitch_email/$(SQITCH_EMAIL)"/g $(ARG)/sqitch.conf
	sed -i "s/@@engine_client/$(SQITCH_ENGINE_CLIENT)/g" $(ARG)/sqitch.conf
	sed -i "s/@@engine/$(SQITCH_ENGINE)/g" $(ARG)/sqitch.conf
	sed -i "s/@@db_user/$(DB_USER)"/g $(ARG)/sqitch.conf
	sed -i "s/@@db_password/$(PGPASSWORD)"/g $(ARG)/sqitch.conf
	sed -i "s/@@db_host/$(DB_HOST)"/g $(ARG)/sqitch.conf
	sed -i "s/@@db_port/$(DB_PORT)"/g $(ARG)/sqitch.conf
	cd $(ARG); \
	sqitch init $(DB_NAME)
	echo "sqitch init $(DB_NAME)"

## create postgres database schema
define DB-CREATE
	@if [ -e $(ARG) ]; then \
		podman exec -it $(DB_CONTAINER) createdb -U $(DB_USER) $(DB_NAME); \
		echo "createdb $(DB_NAME)"; \
	else \
		podman exec -it $(DB_CONTAINER) createdb -U $(DB_USER) $(ARG); \
		echo "createdb $(ARG)"; \
	fi
endef

# drop the database (forced)
define DB-DROP
	@if [ -e $(ARG) ]; then \
		podman exec -it $(DB_CONTAINER) dropdb -U $(DB_USER) --force $(DB_NAME); \
		echo "dropped $(DB_NAME)"; \
	else \
		podman exec -it $(DB_CONTAINER) dropdb -U $(DB_USER) --force $(ARG); \
		echo "dropped $(ARG)"; \
	fi
endef

# drop the schema from the database
define DB-DROP-SCHEMA
	@if [ -z "$(1)" ]; then \
		psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS $(DB_SCHEMA) CASCADE"; \
	else \
		psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS $(strip $(1)) CASCADE"; \
	fi
endef

# drop only the sqitch meta-data schema
define DB-DROP-SCHEMA-META
	@if [ -z "$(1)" ]; then \
		psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS sqitch_$(DB_SCHEMA) CASCADE"; \
	else \
		psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS sqitch_$(strip $(1)) CASCADE"; \
	fi
endef

define DB-BACKUP-SCHEMA-META
	@if [ -z "$(1)" ]; then \
		pg_dump -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -n sqitch_$(DB_SCHEMA) -F c -f sqitch_$(DB_SCHEMA).dump; \
	else \
		pg_dump -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -n sqitch_$(strip $(1)) -F c -f sqitch_$(strip $(1)).dump; \
	fi
endef

define DB-RESTORE-SCHEMA-META
	@if [ -z "$(1)" ]; then \
		pg_restore -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -n sqitch_$(DB_SCHEMA) -c -F c -v --single-transaction -x sqitch_$(DB_SCHEMA).dump; \
	else \
		pg_restore -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -n sqitch_$(strip $(1)) -c -F c -v --single-transaction -x sqitch_$(strip $(1)).dump; \
	fi
endef

# remove all the schema directory
define DB-REMOVE-SCHEMA
	@if [ -d "$(ARG)" ]; then \
		echo "schema $(ARG) removed"; \
		rm -rf $(ARG); \
	else \
		echo "schema $(ARG) not found."; \
	fi
endef

define DB-CREATE-SCHEMA
	@if [ -d "$(ARG)" ]; then \
		echo "schema found: $(ARG)"; \
		$(MAKE) db-create-schema ARG=$(ARG); \
	elif [ -n "$(ARG)" ]; then \
		echo "schema created found: $(ARG)"; \
		mkdir -p $(ARG); \
		$(MAKE) db-create-schema ARG=$(ARG); \
	else \
		echo "schema created found: $(DB_SCHEMA_DEFAULT)"; \
		mkdir -p $(DB_SCHEMA_DEFAULT); \
		$(MAKE) db-create-schema ARG=$(DB_SCHEMA_DEFAULT); \
	fi
endef

# dump table
# pg_dump -h your_host -U your_user -d your_database -t your_table --schema-only > table_schema.sql
# pg_dump -h localhost -U postgres -d underground -t circle.names --schema-only > circle_schema.sql

# dump schema
# pg_dump -h your_host -U your_user -d your_database -n schema1,schema2 -Fc -f your_backup.dump
# pg_dump -h localhost -U postgres -d underground -n circle -f circle.sql
