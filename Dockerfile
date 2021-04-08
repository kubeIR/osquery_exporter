FROM centos:8

LABEL maintainer=prateeknischal \
    version=0.0.1-alpha \
    project=kubeir

RUN yum update -y && yum install -y git openssl wget

ENV REPO="https://github.com/prateeknischal/osquery_exporter/releases/download"

RUN  wget "https://pkg.osquery.io/rpm/osquery-4.7.0-1.linux.x86_64.rpm" && \
    rpm -i osquery-4.7.0-1.linux.x86_64.rpm && \
    echo "{}" > /etc/osquery/osquery.conf

ARG VERSION=0.2.0-alpha

RUN wget $REPO/$VERSION/osquery_exporter_${VERSION}_linux_amd64.tar.gz && \
    mkdir -p /etc/osquery && \
    tar -C /etc/osquery -xf osquery_exporter_${VERSION}_linux_amd64.tar.gz && \
    mv /etc/osquery/osquery_exporter /etc/osquery/osquery_exporter.ext && \
    chown root:root /etc/osquery/osquery_exporter.ext && \
    chmod 500 /etc/osquery/osquery_exporter.ext && \
    echo "/etc/osquery/osquery_exporter.ext" > /etc/osquery/extensions.load

EXPOSE 5000

CMD ["osqueryd"]

