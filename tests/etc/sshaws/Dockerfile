FROM golang:alpine

RUN apk add make curl

RUN apk --no-cache add ca-certificates &&\
    rm -rf /var/cache/apk/*

COPY tests/etc/sshaws/certificate_authority/myCA.pem /usr/local/share/ca-certificates

RUN update-ca-certificates

WORKDIR /opt/sshaws

COPY . /opt/sshaws

COPY tests/etc/sshaws/dummy_credentials/* /root/.aws/
