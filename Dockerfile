FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags '-extldflags "-static"' -o /go-todo-cli .


FROM alpine:latest

WORKDIR /app
COPY --from=builder /go-todo-cli .
EXPOSE 3000

CMD ["./go-todo-cli", "serve"]