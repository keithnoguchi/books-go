# SPDX-License-Identifier: GPL-2.0
.PHONY: all build clean
all: fmt vet build test
build:
	@mkdir -p ./bin && go build -o ./bin ./...
clean:
	@go clean && rm ./bin/*
%:
	@go $@ ./...
