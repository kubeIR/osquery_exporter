FROM golang:1.15.2-alpine3.12

LABEL maintainer=prateeknischal \
    version=0.0.1-alpha \
    project=kubeir

RUN apk add git openssl wget

ENV VERSION=0.0.1-alpha
ENV REPO="https://github.com/prateeknischal/osquery_exporter/releases/download"

RUN wget $REPO/$VERSION/osquery_exporter_${VERSION}_linux_amd64.tar.gz && \
    mkdir -p /opt/osquery_exporter && \
    tar -C /opt/osquery_exporter -xf osquery_exporter_${VERSION}_linux_amd64.tar.gz

RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 5000

WORKDIR /opt/osquery_exporter
CMD ["/opt/osquery_exporter/osquery_exporter", "--socket=/var/osquery/osquery.em"]

