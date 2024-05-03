# @include-pre@
.PHONY: generate setup-sqlc
# @include-pre$

# @target@
generate:  ## generate queries
	rm -rf sql/db
	@sqlc generate -f sqlc/sqlc.yaml > /dev/null 2>&1 || ( \
		sqlc generate -f sqlc/sqlc.yaml; \
		printf "❌ $(RED)Error Running SQLC$(NC)\n"; \
		printf "    🔹 $(YELLOW)Make sure database migrations are up-to-date.$(NC)\n"; \
		printf "    🔹 $(YELLOW)Make sure schema path in sqlc.yaml file is correct.$(NC)\n"; \
		printf "    🔹 $(YELLOW)Make sure sql queries are correct.$(NC)\n"; \
		exit 1; \
	)
	@echo "✅ sqlc generate successful"

sqlc-install: ## install or upgrade sqlc
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest > /dev/null 2>&1 || (echo "❌ Error: Failed to install sqlc"; exit 1)
	@echo "✅ sqlc installed"
# @target$
