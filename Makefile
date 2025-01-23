# Makefile

# Define variables
GQLGEN_VERSION ?= v0.17.$(v)
GQLGEN_PACKAGE = github.com/99designs/gqlgen@$(GQLGEN_VERSION)
GQLGEN_GENERATE = github.com/99designs/gqlgen

.PHONY: all install-gqlgen generate clean

generate: install-gqlgen run

install-gqlgen:
	go get $(GQLGEN_PACKAGE)
run:
	go run $(GQLGEN_GENERATE) generate

# Clean up Go cache (optional)
clean:
	go clean -modcache
