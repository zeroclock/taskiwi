FROM golang:1.12.17

WORKDIR /go/src

ENV GO111MODULE=on

COPY . /go/src

RUN apt-get update \
    && apt-get install git \
    && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
