# SPDX-License-Identifier: GPL-2.0
MODS	:= book concurrency
.PHONY:	all build clean run bench
all:	fmt vet build test
build:
	@for mod in $(MODS); do \
		cd $${mod} && mkdir -p ./bin && go build -o ./bin ./... && cd ..; \
	done
clean:
	@for mod in $(MODS); do \
		cd $${mod} && go clean && rm -f ./bin/* && cd ..; \
	done
run: build
	@for mod in $(MODS); do \
		cd $${mod} && for cmd in ./bin/*; do $${cmd}; done && cd ..; \
	done
test:
	@for mod in $(MODS); do \
		cd $${mod} && go test -cover ./...; cd ..; \
	done
bench:
	@for mod in $(MODS); do \
		cd $${mod} && go test -benchtime=5s -bench=. ./... && cd ..; \
	done
%:
	@for mod in $(MODS); do cd $${mod} && go $@ ./... && cd ..; done
