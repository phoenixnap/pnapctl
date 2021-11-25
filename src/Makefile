# Build automation directives.
# Content of this file is heavely inspired by https://github.com/vincentbernat/hellogopher

MODULE             = $(shell env GO111MODULE=on $(GO) list -m)
CLI_NAME           = pnapctl
DATE              ?= $(shell date +%FT%T%z)
VERSION           ?= $(shell git describe --tags --always --match=v*.*.* 2> /dev/null || echo v0)
LATEST_STABLE_TAG := $(shell git tag -l "v*.*.*" --sort=-v:refname | awk '!/rc/' | head -n 1)
REVISION          := $(shell git rev-parse --short=8 HEAD || echo unknown)
BRANCH            := $(shell git show-ref | grep "$(REVISION)" | grep -v HEAD | awk '{print $$2}' | sed 's|refs/remotes/origin/||' | sed 's|refs/heads/||' | sort | head -n 1)
PKGS               = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))

BUILD_PLATFORMS  = linux/amd64 darwin/amd64 windows/amd64
ENVIRONMENT_NAME = dev

TESTPKGS = $(shell env GO111MODULE=on $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))

GO      = go
TIMEOUT = 15

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

ARTIFACT_FOLDER      = $(BUILD)/$(ENVIRONMENT_NAME)
ARTIFACT_DIST_FOLDER = $(ARTIFACT_FOLDER)/dist

export GO111MODULE=on

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOX = $(BIN)/gox
$(BIN)/gox: PACKAGE = github.com/mitchellh/gox

GO_JUNIT_REPORT = $(BIN)/go-junit-report
$(BIN)/go-junit-report: PACKAGE = github.com/jstemmer/go-junit-report

# Binaries

.PHONY: build
build: $(GOX) ; $(info $(M) building executable…) @ ## Build cross compilation binaries ready for deployment
	$Q $(GOX) -osarch="$(BUILD_PLATFORMS)" -output="$(ARTIFACT_FOLDER)/$(CLI_NAME)-{{.OS}}-{{.Arch}}" -tags="$(ENVIRONMENT_NAME)" \
		-tags $(ENVIRONMENT_NAME) \
		-ldflags '-X $(MODULE)/commands/version.Version=$(VERSION) -X $(MODULE)/commands/version.BuildDate=$(DATE) -X $(MODULE)/commands/version.BuildCommit=$(REVISION)'

.PHONY: build-simple
build-simple: $(BIN) ; $(info $(M) building executable…) @ ## Simple build process used for local development
	$Q $(GO) build \
		-tags $(ENVIRONMENT_NAME) \
		-ldflags '-X $(MODULE)/commands/version.Version=$(VERSION) -X $(MODULE)/commands/version.BuildDate=$(DATE) -X $(MODULE)/commands/version.BuildCommit=$(REVISION)' \
		-o $(BIN)/$(basename $(CLI_NAME)) main.go

.PHONY: pack
pack: ; $(info $(M) packing executables…) @ ## Pack generated cross compilation binaries
	mkdir $(ARTIFACT_DIST_FOLDER) && \
	tar -czf $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-darwin-amd64.tar.gz --transform='flags=r;s|$(CLI_NAME)-darwin-amd64|$(CLI_NAME)|' -C $(ARTIFACT_FOLDER) $(CLI_NAME)-darwin-amd64 && \
	tar -czf $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-linux-amd64.tar.gz --transform='flags=r;s|$(CLI_NAME)-linux-amd64|$(CLI_NAME)|' -C $(ARTIFACT_FOLDER) $(CLI_NAME)-linux-amd64 && \
	mv $(ARTIFACT_FOLDER)/$(CLI_NAME)-windows-amd64.exe $(ARTIFACT_FOLDER)/$(CLI_NAME).exe && zip -j $(ARTIFACT_DIST_FOLDER)/$(CLI_NAME)-windows-amd64.zip $(ARTIFACT_FOLDER)/$(CLI_NAME).exe

build-and-pack: ; @ ## Build cross compilation binaries ready for deployment and pack them for distibution
	make version
	make clean-build
	make build
	make pack

# Tests

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests test-tparse
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test -vet=off -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)
test-tparse: ; $(info $(M) running $(NAME:%=% )tests...) @
	$Q $(GO) test -vet=off -timeout $(TIMEOUT)s -json -cover $(TESTPKGS) | tparse -all 

.PHONY: test-coverage test-coverage-tools test-coverage-show
test-coverage-show: ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	 $Q $(GO) test -vet=off -coverprofile cover.out $(ARGS) $(TESTPKGS) && go tool cover -html=cover.out && rm cover.out
test-coverage-tools: | $(GO_JUNIT_REPORT)
test-coverage: test-coverage-tools ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)
	$Q $(GO) test \
		-v \
		-vet=off \
		-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $(TESTPKGS) | \
					grep '^$(MODULE)/' | \
					tr '\n' ',' | sed 's/,$$//') \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile="$(COVERAGE_PROFILE)" $(TESTPKGS) 2>&1 | \
		$(GO_JUNIT_REPORT) > $(COVERAGE_DIR)/report.xml

# Misc

.PHONY: clean-build
clean-build: ; $(info $(M) cleaning build directory…)	@ ## Cleanup build directory
	@rm -rf $(ARTIFACT_FOLDER)


.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo Current version: $(VERSION)
	@echo Current revision: $(REVISION)
	@echo Current branch: $(BRANCH)
	@echo Current date: $(DATE)
	@echo Build platforms: $(BUILD_PLATFORMS)
	@echo Latest stable tag: $(LATEST_STABLE_TAG)
