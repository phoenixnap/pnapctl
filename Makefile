SRC                          = $(CURDIR)/src
DOCS                         = $(CURDIR)/docs
DOCS_GENERATION              = $(DOCS)/generation
COMPONENT_TESTS              = $(CURDIR)/component-tests
TEST_RESULTS_DIR             = $(CURDIR)/out
COMPONENT_TEST_RESULTS_DIR   = $(TEST_RESULTS_DIR)/component-tests

COMPONENT_TEST_SUPPORT_LIB   = $(COMPONENT_TESTS)/support/lib
BATS_SUPPORT                 = $(COMPONENT_TEST_SUPPORT_LIB)/bats-support
BATS_SUPPORT_LOADER          = $(BATS_SUPPORT)/load.bash
BATS_ASSERT                  = $(COMPONENT_TEST_SUPPORT_LIB)/bats-assert
BATS_ASSERT_LOADER           = $(BATS_ASSERT)/load.bash
BATS_SUPPORT_VERSION         = v0.3.0
BATS_ASSERT_VERSION          = v2.0.0

export BIN                   = $(CURDIR)/bin
export BUILD                 = $(CURDIR)/build
export UNIT_TEST_RESULTS_DIR = $(TEST_RESULTS_DIR)/unit-tests

BATS = bats
GIT  = git

MAKE_FLAGS = -s

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# Directories 

$(COMPONENT_TEST_RESULTS_DIR):
	$Q mkdir -p $(COMPONENT_TEST_RESULTS_DIR)

# Dependencies

$(BATS_SUPPORT_LOADER) $(BATS_ASSERT_LOADER): ; $(info $(M) fetching bats libraries...)
	$Q $(GIT) submodule update --init ;

bats-support-verify-version: $(BATS_SUPPORT_LOADER) ; $(info $(M) verifying bats-support version '$(BATS_SUPPORT_VERSION)'...)
	$Q cd $(BATS_SUPPORT) && git checkout $(BATS_SUPPORT_VERSION)

bats-assert-verify-version: $(BATS_ASSERT_LOADER) ; $(info $(M) verifying bats-assert version '$(BATS_ASSERT_VERSION)'...)
	$Q cd $(BATS_ASSERT) && git checkout $(BATS_ASSERT_VERSION)

# Binaries

.PHONY: build
build build-simple pack build-and-pack: generate-docs
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

# Tests

.PHONY:
generate-all-mocks generate-all-mock-clients generate-all-mock-utils test-bench test-short test-verbose test-race check test tests test-tparse test-coverage-show test-coverage:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

.PHONY: component-tests
component-tests: build-simple bats-support-verify-version bats-assert-verify-version $(COMPONENT_TEST_RESULTS_DIR) ; $(info $(M) running component tests…) @ ## Run Component Tests
	$Q cd $(COMPONENT_TESTS) && $(BATS) --report-formatter junit -o $(TEST_RESULTS_DIR)/component-tests/ .

# Misc

.PHONY: generate-docs
generate-docs: 
	$Q $(MAKE) $(MAKE_FLAGS) -C $(DOCS_GENERATION) $@

.PHONY:
version:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup bin, build, and test result directories
	@rm -rf $(BIN)
	@rm -rf $(BUILD)
	@rm -rf $(TEST_RESULTS_DIR)
	@rm -rf $(DOCS_GENERATION)/bin
	@find $(DOCS) -name '*.md' -type f -delete

.PHONY: help
help:
	$Q @grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@
	$Q $(MAKE) $(MAKE_FLAGS) -C $(DOCS_GENERATION) $@
