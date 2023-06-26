FROM golang:latest

WORKDIR /app

COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
COPY ./app/ ./
RUN go install github.com/cosmtrek/air@latest

CMD ["go", "run", "cmd/main.go"]