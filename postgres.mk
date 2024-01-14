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
	sed -i "s/@@db_password/$(DB_PASSWORD)"/g $(ARG)/sqitch.conf
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

define DB-DROP
	@if [ -e $(ARG) ]; then \
		podman exec -it $(DB_CONTAINER) dropdb -U $(DB_USER) --force $(DB_NAME); \
		echo "dropped $(DB_NAME)"; \
	else \
		podman exec -it $(DB_CONTAINER) dropdb -U $(DB_USER) --force $(ARG); \
		echo "dropped $(ARG)"; \
	fi
endef

define DB-DROP-SCHEMA
    psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA IF EXISTS sqitch CASCADE"
endef

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
