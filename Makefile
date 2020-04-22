# SPDX-License-Identifier: GPL-2.0
BOOKS	:= book
BOOKS	+= concurrency
BOOKS	+= distributed
PROTO	:= distributed

.PHONY:	all mod proto build clean run bench
all:	fmt vet mod proto build test
mod:
	@for book in $(PROTO); do $$(cd $${book} && go mod download); done
proto: mod
	@go get -u github.com/gogo/protobuf/protoc-gen-gogo
	@for book in $(PROTO); do \
		$$(cd $${book} && protoc ch02/api/v1/*.proto \
			--gogo_out=Mgogoproto/gogo.proto=github.com/gogo/protobuf/proto:. \
			--proto_path=$$(go list -f '{{ .Dir }}' -m github.com/gogo/protobuf) \
			--proto_path=. ); \
	done
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
