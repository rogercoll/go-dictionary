VERSION=$(shell cat VERSION)
APP_NAME=mydict

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container
	podman build -t $(APP_NAME) .

deploy: ## Build the container
	podman run $(APP_NAME) 
