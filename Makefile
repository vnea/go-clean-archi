# This is an auto documented Makefile. For more information see the following article
# @see http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

SHELL := /bin/bash
OS := $(shell uname)

.SHELLFLAGS = -ec
.SILENT:
MAKEFLAGS += --silent
.ONESHELL:

.EXPORT_ALL_VARIABLES:

.DEFAULT_GOAL: help

.PHONY: help ## 🆘 This will list all available targets with their documentation.
help:
	echo "❓ Use \`make <target>' where <target> is one of:"
	grep -E '^\.PHONY: [a-zA-Z0-9_-]+ .*?##' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = "(: |##)"}; {printf "\033[36m%-30s\033[0m %s\n", $$2, $$3}'

.PHONY: build ## 🔨 To build the application
build:
	go build -o build/go-clean-archi

.PHONY: run ## 🛬 To run the application
run:
	go run .

.PHONY: test ## 📖 To run the tests
test:
	go test ./...

.PHONY: get-albums ## 🔖 To get the albums
get-albums:
	curl --silent --show-error http://localhost:8080/albums | jq
