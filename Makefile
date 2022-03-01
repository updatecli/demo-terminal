
.PHONY: quickstart
quickstart: ## Run Updatecli Quickstart Demo
	cd quickstart && clear && go run main.go -0

.PHONY: help
help: ## Show this Makefile's help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

