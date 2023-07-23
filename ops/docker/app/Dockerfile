FROM golang:latest

WORKDIR /data
COPY ./ ./

WORKDIR /data/app
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["go", "run", "cmd/main.go"]