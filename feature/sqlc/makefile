# @include-pre@
.PHONY: generate setup-sqlc
# @include-pre$

# @target@
generate: ## generate queries
	rm -rf db
	sqlc generate

setup-sqlc: ## install or upgrade sqlc
	@if [ "$(OS)" = "Darwin" ]; then \
		brew reinstall sqlc; \
	fi
# @target$
