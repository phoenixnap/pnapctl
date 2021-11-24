SRC                         = $(CURDIR)/src
TEST_RESULTS_DIR            = $(CURDIR)/test
export BIN                  = $(CURDIR)/bin
export BUILD                = $(CURDIR)/build
export COVERAGE_DIR         = $(TEST_RESULTS_DIR)/coverage

# Binaries

.PHONY: build
build build-simple pack build-and-pack:
	$Q $(MAKE) -C $(SRC) $@

# Tests

.PHONY:
test-bench test-short test-verbose test-race check test tests test-coverage-show test-coverage:
	$Q $(MAKE) -C $(SRC) $@

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf $(BUILD)
	@rm -rf $(TEST_RESULTS_DIR)

.PHONY: help
help:
	$Q $(MAKE) -C $(SRC) $@