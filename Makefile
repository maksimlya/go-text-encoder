BUILD_MODE?=c-shared
OUTPUT_DIR?=output
GO_BINARY?=go
BINDING_NAME?=libencoder_bridge
BINDING_FILE?=$(BINDING_NAME).so
BINDING_ARGS?=
BINDING_OUTPUT?=$(OUTPUT_DIR)/binding
EXTRA_LD_FLAGS?=

default: fmt test

clean:
	rm -rf output

deps:
	go mod download

binding: deps
	mkdir -p $(BINDING_OUTPUT)
	$(GO_BINARY) build -ldflags="-w -s $(EXTRA_LD_FLAGS)" -o $(BINDING_OUTPUT)/$(BINDING_FILE) -buildmode=$(BUILD_MODE) $(BINDING_ARGS) binding/main.go

include Makefile.android
