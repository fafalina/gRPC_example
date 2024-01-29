FROM golang:alpine AS builder

WORKDIR /src

COPY . /src
RUN cd /src
RUN go mod download
RUN go build -o app

FROM alpine
COPY --from=builder /src/app /
ENTRYPOINT ["/app"]

EXPOSE 50051