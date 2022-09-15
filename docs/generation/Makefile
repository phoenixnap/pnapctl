GO   = go
BIN  = bin
DOCS = $(BIN)/docgenerator

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

export GO111MODULE=on

.PHONY: generate-docs 
generate-docs: remove-outdated-docs build-docgenerator run-docgenerator

remove-outdated-docs: ; $(info $(M) removing outdated docs...) @
	$Q rm -f ../*.md

build-docgenerator: main.go go.mod go.sum; $(info $(M) building docgenerator executable docs…) @
	$Q $(GO) build -o $(DOCS) main.go

run-docgenerator: $(DOCS); $(info $(M) generating docs…) @
	$Q ./$(DOCS)

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'