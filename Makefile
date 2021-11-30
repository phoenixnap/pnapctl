SRC                          = $(CURDIR)/src
COMPONENT_TESTS              = $(CURDIR)/component-tests
TEST_RESULTS_DIR             = $(CURDIR)/out
COMPONENT_TEST_RESULTS_DIR   = $(TEST_RESULTS_DIR)/component-tests

COMPONENT_TEST_SUPPORT_LIB   = $(COMPONENT_TESTS)/support/lib
BATS_SUPPORT_LOADER          = $(COMPONENT_TEST_SUPPORT_LIB)/bats-support/load/bash
BATS_ASSERT_LOADER           = $(COMPONENT_TEST_SUPPORT_LIB)/bats-assert/load.bash

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
	$Q cd $(COMPONENT_TEST_SUPPORT_LIB)/bats-support && git checkout v0.3.0 ;
	$Q cd $(COMPONENT_TEST_SUPPORT_LIB)/bats-assert && git checkout v2.0.0

# Binaries

.PHONY: build
build build-simple pack build-and-pack:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

# Tests

.PHONY:
test-bench test-short test-verbose test-race check test tests test-coverage-show test-coverage:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

.PHONY: component-tests
component-tests: build-simple $(BATS_SUPPORT_LOADER) $(BATS_ASSERT_LOADER) $(COMPONENT_TEST_RESULTS_DIR) ; $(info $(M) running component tests…) @ ## Run Component Tests
	$Q cd $(COMPONENT_TESTS) && $(BATS) --report-formatter junit -o $(TEST_RESULTS_DIR)/component-tests/ .

# Misc

.PHONY:
version:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup bin, build, and test result directories
	@rm -rf $(BIN)
	@rm -rf $(BUILD)
	@rm -rf $(TEST_RESULTS_DIR)

.PHONY: help
help:
	$Q @grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@