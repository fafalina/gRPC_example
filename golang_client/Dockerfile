FROM golang:alpine AS builder

WORKDIR /client

COPY . .
RUN go mod download
RUN go build -o app

FROM alpine:latest
COPY --from=builder /client/app /

ENTRYPOINT ["/app"]
