GOFLAGS := GOOS=linux CGO_ENABLED=0
GO := $(GOFLAGS) go

# with build flags for static linking
GOBUILD := $(GO) build -ldflags '-extldflags "-static"' -a

SRC_PATH := ./cmd/app
BIN_PATH := ./build/app
DOCKER_IMAGE := backend-assignment

binary: ## build static binary from Go source
	@echo "exporting binary at path" $(BIN_PATH)
	@$(GOBUILD) -o $(BIN_PATH) $(SRC_PATH)
	@strip $(BIN_PATH)

clean: ## remove all the built binaries and docker image
	@rm -rf build/
	-@docker rmi $(DOCKER_IMAGE)

docker: ## build the docker image (compiles the binaries in the container)
	@docker build -t $(DOCKER_IMAGE) .

generate: ## generate everything in "./pkg/**/*/generate.go"
	@$(GO) generate ./...

run: ## run the binary locally
	@$(GO) run $(SRC_PATH) --debug

tests: ## run unit tests
	@$(GO) test ./pkg/...

help: ## insert recursive message https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: binary clean docker run tests help
.DEFAULT_GOAL := help
