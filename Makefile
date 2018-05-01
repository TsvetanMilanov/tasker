SHELL               := /bin/bash
BUILD_DIR           := .build
STAGE               ?= dev
ENV_FILE_NAME       := $(STAGE)-env.list
SERVICES_DIR        := src/services
DISCOVERED_SERVICES := $(shell ls $(SERVICES_DIR))

GO_NAMESPACE        := github.com/TsvetanMilanov/tasker

GREEN               := \e[32m
NC                  := \e[0m

define find_recursive
$(shell find $1 -iname "$2")
endef

define print
	@echo -e '$(GREEN)$1$(NC)'
endef

define get_service_dir
$(shell echo $(SERVICES_DIR)/$1)
endef

define get_service_go_package
$(shell echo $(GO_NAMESPACE)/$(SERVICES_DIR)/$1)
endef

.PRECIOUS: \
	$(BUILD_DIR)/vendor-% \
	$(BUILD_DIR)/build-% \
	$(BUILD_DIR)/deploy-%

.PHONY: \
	build \
	build/% \
	deploy \
	deploy/% \
	vendor-update \
	vendor-update/% \
	clean

.SECONDEXPANSION:
$(BUILD_DIR)/vendor-%: $$(call get_service_dir,%)/Gopkg.toml $$(call get_service_dir,%)/Gopkg.lock
	$(call print,Installing deps for service $*...)
	pushd $(call get_service_dir,$*) && \
	dep ensure && \
	popd
	@touch $@

.SECONDEXPANSION:
$(BUILD_DIR)/build-%: $$(call find_recursive,$(SERVICES_DIR)/%,*.go) $$(BUILD_DIR)/vendor-%
	$(call print,Building service $*...)
	GOOS=linux go build -ldflags="-s -w" \
		-o $(CURDIR)/$(call get_service_dir,$*)/bin/$* \
		$(call get_service_go_package,$*)
	@touch $@

.SECONDEXPANSION:
$(BUILD_DIR)/deploy-%: $(BUILD_DIR)/build-% $$(call get_service_dir,%)/serverless.yml $(CURDIR)/infra/scripts/sls-common.js
	$(call print,Deploying service $*...)
	source $(ENV_FILE_NAME); pushd $(call get_service_dir,$*) && \
	serverless deploy --stage $(STAGE) && \
	popd
	@touch $@

build/%: $(BUILD_DIR)/build-%
	$(call print,Service $* successfully built)

build: $(patsubst %,build/%,$(DISCOVERED_SERVICES))
	$(call print,All services successfully built)

deploy/%: $(BUILD_DIR)/deploy-%
	$(call print,Service $* successfully deployed)

deploy: $(patsubst %,deploy/%,$(DISCOVERED_SERVICES))
	$(call print,All services successfully deployed)

vendor-update/%:
	$(call print,Updating vendors for service $*...)
	pushd $(call get_service_dir,$*) && \
	dep ensure -update && \
	popd

vendor-update: $(patsubst %,vendor-update/%,$(DISCOVERED_SERVICES))
	$(call print,All vendors successfully updated)

clean:
	rm -rf $(BUILD_DIR)/*
