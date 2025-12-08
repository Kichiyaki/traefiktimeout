export GO111MODULE=on

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBIN=$(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN := $(shell go env GOPATH)/bin
endif
OSARCH=$(shell uname -m)
GOLANGCI_LINT_PATH=$(GOBIN)/golangci-lint
YAEGI_PATH=$(GOBIN)/yaegi

.PHONY: install-git-hooks
install-git-hooks:
	@echo "Installing git hooks..."
	pre-commit install --hook-type pre-commit
	pre-commit install --hook-type commit-msg

.PHONY: install-golangci-lint
install-golangci-lint:
	@echo "Installing github.com/traefik/yaegi..."
	@(test -f $(GOLANGCI_LINT_PATH) && echo "github.com/traefik/yaegi is already installed. Skipping...") || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v2.7.2

.PHONY: install-yaegi
install-yaegi:
	@echo "Installing github.com/traefik/yaegi..."
	@(test -f $(YAEGI_PATH) && echo "github.com/traefik/yaegi is already installed. Skipping...") || curl -sfL https://raw.githubusercontent.com/traefik/yaegi/master/install.sh | bash -s -- -b $(GOBIN) v0.16.1

.PHONY: install-tools
install-tools: install-golangci-lint install-yaegi

.PHONY: install
install: install-tools install-git-hooks

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -cover --race ./...

.PHONY: yaegi-test
yaegi-test:
	yaegi test -v .

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: clean
clean:
	rm -rf ./vendor
