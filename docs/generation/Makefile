GO   = go
BIN  = bin
DOCS = $(BIN)/docs

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

$(DOCS): ; $(info $(M) building docs executable…) 
	$Q $(GO) build -o $(DOCS) main.go

.PHONY: generate-docs 
generate-docs: $(DOCS) ; $(info $(M) generating docs…) @ ## Generate Commands documentation
	$Q ./$(DOCS)

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'