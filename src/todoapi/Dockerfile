FROM golang:1.14 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/todoapi

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /todoapi ./main.go

FROM alpine:latest

COPY --from=builder /todoapi /app/todoapi

ENV PORT=3000

ENTRYPOINT ./app/todoapi