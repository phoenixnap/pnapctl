SRC                         = $(CURDIR)/src
export BIN                  = $(CURDIR)/bin
export BUILD                = $(CURDIR)/build
export COVERAGE_DIR         = $(CURDIR)/test

# Binaries

.PHONY: build
build build-simple pack build-and-pack:
	$Q $(MAKE) -C $(SRC) $@

# Tests

.PHONY:
generate-mock test-bench test-short test-verbose test-race check test tests test-coverage-show test-coverage:
	$Q $(MAKE) -C $(SRC) $@

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf $(BUILD)
	@rm -rf $(COVERAGE_DIR)

.PHONY: help
help:
	$Q $(MAKE) -C $(SRC) $@