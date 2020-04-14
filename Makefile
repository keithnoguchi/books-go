# SPDX-License-Identifier: GPL-2.0
BOOKS	:= book
BOOKS	+= concurrency
BOOKS	+= distributed

.PHONY:	all build clean run bench
all:	fmt vet build test
build:
	@for book in $(BOOKS); do \
		cd $${book} && mkdir -p ./bin && go build -o ./bin ./... && cd ..; \
	done
clean:
	@for book in $(BOOKS); do \
		cd $${book} && go clean && rm -f ./bin/* && cd ..; \
	done
run: build
	@for book in $(BOOKS); do \
		cd $${book} && for cmd in ./bin/*; do $${cmd}; done && cd ..; \
	done
test:
	@for book in $(BOOKS); do \
		cd $${book} && go test -cover ./...; cd ..; \
	done
bench:
	@for book in $(BOOKS); do \
		cd $${book} && go test -benchtime=5s -cpu=1 -bench=. ./... && cd ..; \
	done
%:
	@for book in $(BOOKS); do cd $${book} && go $@ ./... && cd ..; done
