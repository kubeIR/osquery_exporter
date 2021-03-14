FROM golang:1.15.2-alpine3.12

LABEL maintainer=prateeknischal \
    version=0.1.0 \
    project=kubeir

RUN apk add git make openssl zip gcc musl-dev

