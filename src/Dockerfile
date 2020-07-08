FROM golang:1.14.4

WORKDIR /go/src

ENV GO111MODULE=on

COPY . /go/src

RUN apt-get update \
    && apt-get install -y git python jq curl \
    && curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get update && apt-get install -y nodejs \
    && npm install yarn -g \
    && wget https://github.com/go-task/task/releases/download/v2.8.1/task_linux_amd64.deb \
    && dpkg -i task_linux_amd64.deb \
    && rm task_linux_amd64.deb

EXPOSE 8080

CMD ["./startup.sh"]
