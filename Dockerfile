FROM golang:stretch
MAINTAINER Positron <positron@jarm.com>

ENV GNODE_IP "111.222.111.222"
ENV GNODE_PORT "26968"
ENV GNODE_WEB "26969"

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["go-wrapper", "run"]

COPY . /go/src/app
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install
