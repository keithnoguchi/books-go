# SPDX-License-Identifier: GPL-2.0
FROM archlinux/base
RUN pacman -Sy --noconfirm make libffi gcc go git protobuf
RUN go get -u github.com/gogo/protobuf/protoc-gen-gogo
ENV GOPATH=/root/go
ENV PATH=$GOPATH/bin:$PATH
ADD . /root/build
WORKDIR /root/build
CMD ["make"]
