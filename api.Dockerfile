FROM golang:latest-alpine

WORKDIR /app

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

RUN go get github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint,.sh

RUN chmod +rx ./entrypoint.sh

ENTRYPOINT [ "sh","./entrypoint.sh" ]