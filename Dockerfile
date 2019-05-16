FROM golang:1.12 AS builder
WORKDIR /gorum

COPY go.mod ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN CGO_ENABLED=0 go build -o /go/bin/server -v ./cmd/server

FROM alpine:3.9 AS final

COPY --from=builder /go/bin/server /bin/server

COPY .docker/docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]