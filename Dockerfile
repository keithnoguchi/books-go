# SPDX-License-Identifier: GPL-2.0
FROM archlinux/base
RUN pacman -Sy --noconfirm make libffi gcc go git protobuf
ENV GOPATH=/root/go
ENV PATH=$GOPATH/bin:$PATH
ADD . /root/build
WORKDIR /root/build
CMD ["make"]
