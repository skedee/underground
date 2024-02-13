# @target@
generate: ## generate queries
	sqlc generate

setup-sqlc: ## install or upgrade sqlc
	@if [ "$(OS)" = "Darwin" ]; then \
		brew reinstall sqlc; \
	fi
# @target$
