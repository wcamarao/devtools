# Module
GO_MOD = $(shell head -n 1 go.mod|sed "s/^module //g")
MOD_NAME = $(shell basename $(GO_MOD))

# Database
export DB_HOST ?= localhost
export DB_USER ?= codelab
export DB_PASS ?= codelab
export DB_NAME ?= $(MOD_NAME)
export DB_SSLMODE ?= disable
DB_BASE_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST)
DB_URL = $(DB_BASE_URL)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
CREATE_DB = 'create database "$(DB_NAME)"'
DROP_DB = 'drop database "$(DB_NAME)"'
DB_SCHEMA = db/schema.sql

# Misc
HIGHLIGHT = @printf "\n\033[36m>> $1\033[0m\n"

default: help

help:
	@echo "Usage: make <TARGET> [OPTS=opts]\n\nTargets:"
	@grep -E "^[\. a-zA-Z_-]+:.*?## .*$$" $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' |sort

test: ## Run unit tests
	$(call HIGHLIGHT,test)
	@go test -count=1 ./...

db.create: ## Create database
	$(call HIGHLIGHT,db.create)
	@echo ">" $(CREATE_DB)
	@psql $(DB_BASE_URL) -c $(CREATE_DB)
	@echo ">" $(DB_SCHEMA)
	@psql $(DB_URL) -f $(DB_SCHEMA)

db.drop: ## Drop database
	$(call HIGHLIGHT,db.drop)
	@echo ">" $(DROP_DB)
	@psql $(DB_BASE_URL) -c $(DROP_DB)
