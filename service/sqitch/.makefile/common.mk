# Ascii colors
BLUE = \033[94m
GREEN = \033[92m
NC = \033[0m
PURPLE = \033[95m
RED = \033[91m
YELLOW = \033[93m
CYAN = \033[36m

.PHONY: BLUE GREEN NC PURPLE RED YELLOW CYAN FOREACH PRINT_MENU

define FOREACH
	for DIR in $(1); do \
		$(MAKE) -C ../$$DIR $(2) --no-print-directory; \
	done
endef

define PRINT_MENU
	if [ -n "$(1)" ]; then \
		printf "📁 $(GREEN)$(APP_NAME)$(NC)\n"; \
		grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "🔹 $(YELLOW)%-30s$(NC) %s\n", $$1, $$2}'; \
		echo ""; \
	else \
		printf "📂 $(GREEN)$(APP_NAME)$(NC)\n"; \
		grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "🔹 $(YELLOW)%-30s$(NC) %s\n", $$1, $$2}'; \
		echo ""; \
	fi;
endef
