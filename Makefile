MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
CLI_NAME = pnapctl
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell it describe --tags --always --match=v*.*.* 2> /dev/null || echo v0)
LATEST_STABLE_TAG := $(shell git tag -l "v*.*.*" --sort=-v:refname | awk '!/rc/' | head -n 1)
REVISION := $(shell git rev-parse --short=8 HEAD || echo unknown)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
BUILD_PLATFORMS = linux/amd64 darwin/amd64 windows/amd64
TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))
BIN      = $(CURDIR)/bin

GO      = go
GOX     = gox
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

ARTIFACT_FOLDER = build
ARTIFACT_DIST_FOLDER = dist

export GO111MODULE=on

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = $(BIN)/golint
$(BIN)/golint: PACKAGE=golang.org/x/lint/golint

GOX = $(BIN)/gox
$(BIN)/gox: PACKAGE = github.com/mitchellh/gox

GO_JUNIT_REPORT = $(BIN)/go-junit-report
$(BIN)/go-junit-report: PACKAGE = github.com/mitchellh/gox

# Binaries

.PHONY: build
build: $(GOX) ; $(info $(M) building executable…) @ ## Build program binaries
	$Q $(GOX) -osarch="$(BUILD_PLATFORMS)" -output="build/$(ENVIRONMENT_NAME)/$(CLI_NAME)-{{.OS}}-{{.Arch}}" -tags="$(ENVIRONMENT_NAME)" \
		-tags $(ENVIRONMENT_NAME) \
		-ldflags '-X $(MODULE)/pnapctl/commands/version.Version=$(VERSION) -X $(MODULE)/pnapctl/commands/version.BuildDate=$(DATE) -X $(MODULE)/pnapctl/commands/version.BuildCommit=$(REVISION)'

.PHONY: build-simple
build-simple: $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
	$Q $(GO) build \
		-tags dev \
		-ldflags '-X $(MODULE)/pnapctl/commands/version.Version=$(VERSION) -X $(MODULE)/pnapctl/commands/version.BuildDate=$(DATE) -X $(MODULE)/pnapctl/commands/version.BuildCommit=$(REVISION)' \
		-o $(BIN)/$(basename $(CLI_NAME)) main.go

build-and-deploy:
	make clean-build
	make build
	cd $(ARTIFACT_FOLDER)/$(ENVIRONMENT_NAME) && \
	mkdir $(ARTIFACT_DIST_FOLDER) && \
	tar -czf $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-darwin-amd64.tar.gz --transform='flags=r;s|$(CLI_NAME)-darwin-amd64|$(CLI_NAME)|' $(CLI_NAME)-darwin-amd64 && \
	tar -czf $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-linux-amd64.tar.gz --transform='flags=r;s|$(CLI_NAME)-linux-amd64|$(CLI_NAME)|' $(CLI_NAME)-linux-amd64 && \
	mv $(CLI_NAME)-windows-amd64.exe $(CLI_NAME).exe && zip $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-windows-amd64.zip $(CLI_NAME).exe

# Tests

# Misc

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf $(ARTIFACT_FOLDER)
	@rm -rf test/tests.* test/coverage.*

.PHONY: clean-build
clean-build: ; $(info $(M) cleaning build directory…)	@ ## Cleanup build directory
	@rm -rf $(ARTIFACT_FOLDER)

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
	@echo $(MODULE)
