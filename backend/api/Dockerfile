# ---------------TEST----------------
FROM golang:1.22-alpine

WORKDIR /usr/src/api

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify


RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]