# SPDX-License-Identifier: GPL-2.0
MODS	:= book concurrency
.PHONY:	all build clean run
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
%:
	@for mod in $(MODS); do cd $${mod} && go $@ ./... && cd ..; done
