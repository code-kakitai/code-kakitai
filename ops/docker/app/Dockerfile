FROM golang:1.21.1-alpine3.18

WORKDIR /data
COPY ./ ./

WORKDIR /data/app
RUN go mod download

RUN go install go.uber.org/mock/mockgen@v0.3.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.23.0
