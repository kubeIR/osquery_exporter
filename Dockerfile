FROM golang:latest

COPY ./ /app
WORKDIR /app

ARG VERSION="unknown"

RUN make test && \
    make build VERSION="${VERSION}"

FROM centos:latest

ARG VERSION="latest"
ARG OSQUERY_VERSION="4.7.0-1"
ARG OSQUERY_CONFIG="{}"
ARG OSQUERY_FLAGS=""

LABEL maintainer=prateeknischal \
    version=${VERSION}\
    project=github.com/kubeIR

RUN yum update -y && yum install wget -y

RUN wget "https://pkg.osquery.io/rpm/osquery-${OSQUERY_VERSION}.linux.x86_64.rpm" \
        -O osquery.rpm && \
    rpm -i osquery.rpm && \
    echo "${OSQUERY_CONFIG}" > /etc/osquery/osquery.conf && \
    echo "${OSQUERY_FLAGS}" > /etc/osquery/osquery.flags

COPY --from=0 /app/osquery_exporter /etc/osquery/osquery_exporter.ext

RUN chown root:root /etc/osquery/osquery_exporter.ext && \
    chmod 500 /etc/osquery/osquery_exporter.ext && \
    echo "/etc/osquery/osquery_exporter.ext" > /etc/osquery/extensions.load

EXPOSE 5000

CMD ["osqueryd"]
