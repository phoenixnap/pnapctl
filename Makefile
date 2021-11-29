SRC                         = $(CURDIR)/src
COMPONENT_TESTS             = $(CURDIR)/component-tests
TEST_RESULTS_DIR            = $(CURDIR)/test

export BIN                  = $(CURDIR)/bin
export BUILD                = $(CURDIR)/build
export COVERAGE_DIR         = $(TEST_RESULTS_DIR)/coverage

BATS = bats

MAKE_FLAGS = -s

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# Binaries

.PHONY: build
build build-simple pack build-and-pack:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

# Tests

.PHONY:
test-bench test-short test-verbose test-race check test tests test-coverage-show test-coverage:
	$Q $(MAKE) $(MAKE_FLAGS) -C $(SRC) $@

.PHONY: component-tests
component-tests: build-simple ; $(info $(M) running component tests…) @ ## Run Component Tests
	$Q cd $(COMPONENT_TESTS) && $(BATS) .

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