.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY:
build: ## Build the cli and send it to tem
	go build --race -o /tmp/cli
	@ echo ✅ success!


.PHONY:
run: build ## Run everything given a file
	/tmp/cli run -f 316.toml
	@ echo ✅ success!
